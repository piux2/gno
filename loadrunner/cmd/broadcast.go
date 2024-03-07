
package main

import (
	"context"
	"flag"
	"fmt"
	"strconv"
	"sync"

	"github.com/gnolang/gno/tm2/pkg/amino"
	ctypes "github.com/gnolang/gno/tm2/pkg/bft/rpc/core/types"
	"github.com/gnolang/gno/tm2/pkg/commands"

	"github.com/gnolang/gno/tm2/pkg/crypto/keys"


	"github.com/gnolang/gno/tm2/pkg/bft/rpc/client"

)

type broadcastCfg struct{}

func (c *broadcastCfg) RegisterFlags(fs *flag.FlagSet) {}

func newBroadcastCmd() *commands.Command {
	cfg := &broadcastCfg{}
	return commands.NewCommand(
		commands.Metadata{
			Name:       "broadcast",
			ShortUsage: "broadcast [first n-th signers]",
			ShortHelp:  "broadcast txs stored locally",
		},
		cfg,
		func(_ context.Context, args []string) error {
			return broadcast(args)
		},
	)
}

type BroadcastTask struct {

	SignerAddr      string
	keybase         *EagerKeybase
	txDB            *Txbase
}

func (b BroadcastTask) Key() string {
	return b.SignerAddr
}

func (b BroadcastTask) Execute() error {
	info, err := b.keybase.GetByAddress(b.SignerAddr)
	if err != nil {
		return err
	}
	res , err := broadcastTx(info, b.txDB)
	if err != nil {
		return err
	}

	if res.CheckTx.IsErr() {
		return fmt.Errorf("transaction failed %#v\nlog %s", res, res.CheckTx.Log)
	} else if res.DeliverTx.IsErr() {
		return fmt.Errorf("transaction failed %#v\nlog %s", res, res.DeliverTx.Log)
	}
	fmt.Printf("broadcast: %s\n", b.SignerAddr)
	return nil
}

func broadcast(args []string) error {

	var numWorkers int = 500

	kb, err := NewEagerKeybase(RootDir)
	if err != nil {
		return err
	}
	defer kb.Close()

	// Txbase store signed Txs before it is broadcasted.
	txbase, err := NewTxbase(RootDir, defaultTxDBName)
	if err != nil {
		return err
	}
	defer txbase.Close()
	var n int //  first n-th accounts in tx base
	txinfos := txbase.List()
	l:=len(txinfos)
	if len(args) == 0{ // broadcast all all txs in tx base
			n = l
	}else{
		n, err = strconv.Atoi(args[0])
		if err != nil {
			return nil
		}

		if n > l { // max n is all txs in tx base
			n = l
		}
 }
	// TODO:  check remaining balance vs gas fee

	var wg sync.WaitGroup

	tasks := make(chan Task, n)
	results := make(chan Result, n) // Ensure the buffer is large enough to avoid blocking
	// start the worker routine

	// Start workers
	if n < numWorkers {
		numWorkers = n
	}
	for w := 0; w < numWorkers; w++ {
		go worker(w, tasks, &wg, results)
	}
	// start monitor
	go monitor(results, n)

	// jobs
	wg.Add(n)

	for i:=0; i< n ; i++{
		txinfo := txinfos[i]

		addr := txinfo.addr


		t := BroadcastTask{
			SignerAddr:      addr,
			keybase:         &kb,
			txDB:            &txbase,
		}
		tasks <- t
	}

	wg.Wait()
	close(tasks)
	close(results)

	println()
	fmt.Printf("broadcast %d msgs %d txs in %s/data/txs.db\n", n*numMsgs, n, RootDir)
	return nil
}

func broadcastTx(info keys.Info , txbase *Txbase)(*ctypes.ResultBroadcastTxCommit, error){
	// TODO: here we only have 1 tx per txs, try to use async broadcast with muliple tx
  cli := client.NewHTTP(remote, "/websocket")
	txs, err:=(*txbase).GetTxs(info.GetAddress())
	addr :=info.GetAddress()
	if err !=nil{
		return nil, fmt.Errorf("get txs for %s %w\n",addr, err)
	}

 	var bres *ctypes.ResultBroadcastTxCommit

	for _, tx := range txs{

		bz, err := amino.Marshal(tx)
		if err != nil {
			return nil, fmt.Errorf("remarshaling tx binary bytes %w\n",err)
		}

		bres, err = cli.BroadcastTxCommit(bz)
		if err != nil {
			return nil, fmt.Errorf("broadcasting bytes %w\n", err)
		}



		fmt.Printf("%s GasWanted: %d \tGasUsed: %d\n",addr, bres.DeliverTx.GasWanted,bres.DeliverTx.GasUsed)
    // after it broadcast successfully, we delete txs entires
		(*txbase).Delete(info.GetAddress())


		return bres, nil
	}

  return nil, fmt.Errorf("no transactions found for %s in txs.db\n", info.GetAddress())
}
