package mempool

import (
	"crypto/rand"
	"crypto/sha256"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/gnolang/gno/tm2/pkg/bft/abci/example/kvstore"
	"github.com/gnolang/gno/tm2/pkg/bft/proxy"
	"github.com/gnolang/gno/tm2/pkg/bft/types"
)

func TestCacheRemove(t *testing.T) {
	t.Parallel()

	cache := newMapTxCache(100)
	numTxs := 10
	txs := make([][]byte, numTxs)
	for i := range numTxs {
		// probability of collision is 2**-256
		txBytes := make([]byte, 32)
		rand.Read(txBytes)
		txs[i] = txBytes
		cache.Push(txBytes)
		// make sure its added to both the linked list and the map
		require.Equal(t, i+1, len(cache.map_))
		require.Equal(t, i+1, cache.list.Len())
	}
	for i := range numTxs {
		cache.Remove(txs[i])
		// make sure its removed from both the map and the linked list
		require.Equal(t, numTxs-(i+1), len(cache.map_))
		require.Equal(t, numTxs-(i+1), cache.list.Len())
	}
}

func TestCacheAfterUpdate(t *testing.T) {
	t.Parallel()

	app := kvstore.NewKVStoreApplication()
	cc := proxy.NewLocalClientCreator(app)
	mempool, cleanup := newMempoolWithApp(cc)
	defer cleanup()

	// reAddIndices & txsInCache can have elements > numTxsToCreate
	// also assumes max index is 255 for convenience
	// txs in cache also checks order of elements
	tests := []struct {
		numTxsToCreate int
		updateIndices  []int
		reAddIndices   []int
		txsInCache     []int
	}{
		{1, []int{}, []int{1}, []int{1, 0}},    // adding new txs works
		{2, []int{1}, []int{}, []int{1, 0}},    // update doesn't remove tx from cache
		{2, []int{2}, []int{}, []int{2, 1, 0}}, // update adds new tx to cache
		{2, []int{1}, []int{1}, []int{1, 0}},   // re-adding after update doesn't make dupe
	}
	for tcIndex, tc := range tests {
		for i := range tc.numTxsToCreate {
			tx := types.Tx{byte(i)}
			err := mempool.CheckTx(tx, nil)
			require.NoError(t, err)
		}

		updateTxs := []types.Tx{}
		for _, v := range tc.updateIndices {
			tx := types.Tx{byte(v)}
			updateTxs = append(updateTxs, tx)
		}
		mempool.Update(int64(tcIndex), updateTxs, abciResponses(len(updateTxs), nil), nil, 0)

		for _, v := range tc.reAddIndices {
			tx := types.Tx{byte(v)}
			_ = mempool.CheckTx(tx, nil)
		}

		cache := mempool.cache.(*mapTxCache)
		node := cache.list.Front()
		counter := 0
		for node != nil {
			require.NotEqual(t, len(tc.txsInCache), counter,
				"cache larger than expected on testcase %d", tcIndex)

			nodeVal := node.Value.([sha256.Size]byte)
			expectedBz := sha256.Sum256([]byte{byte(tc.txsInCache[len(tc.txsInCache)-counter-1])})
			// Reference for reading the errors:
			// >>> sha256('\x00').hexdigest()
			// '6e340b9cffb37a989ca544e6bb780a2c78901d3fb33738768511a30617afa01d'
			// >>> sha256('\x01').hexdigest()
			// '4bf5122f344554c53bde2ebb8cd2b7e3d1600ad631c385a5d7cce23c7785459a'
			// >>> sha256('\x02').hexdigest()
			// 'dbc1b4c900ffe48d575b5da5c638040125f65db0fe3e24494b76ea986457d986'

			require.Equal(t, expectedBz, nodeVal, "Equality failed on index %d, tc %d", counter, tcIndex)
			counter++
			node = node.Next()
		}
		require.Equal(t, len(tc.txsInCache), counter,
			"cache smaller than expected on testcase %d", tcIndex)
		mempool.Flush()
	}
}
