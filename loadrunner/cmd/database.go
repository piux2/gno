package main

import (
	"fmt"
	"path/filepath"
	"github.com/gnolang/gno/tm2/pkg/crypto"
	"github.com/gnolang/gno/tm2/pkg/crypto/keys"

	"github.com/gnolang/gno/tm2/pkg/amino"
	tm2db "github.com/gnolang/gno/tm2/pkg/db"
	"github.com/gnolang/gno/tm2/pkg/os"
	"github.com/gnolang/gno/tm2/pkg/std"
)

const (
	dbBackend        = tm2db.GoLevelDBBackend
	defaultKeyDBName = "keys"
	defaultTxDBName  = "txs"
)

type EagerKeybase struct {
	name string
	dir  string
	kb   keys.Keybase
}

// New creates a new instance of a lazy keybase.
func NewEagerKeybase(rootDir string) (EagerKeybase, error) {
	name := defaultKeyDBName
	dir := filepath.Join(rootDir, "data")

	if err := os.EnsureDir(dir, 0o700); err != nil {
		panic(fmt.Sprintf("failed to create Keybase directory: %s", err))
	}
	db, err := tm2db.NewDB(name, dbBackend, dir)
	if err != nil {
		return EagerKeybase{}, err
	}
	kb := keys.NewDBKeybase(db)
	return EagerKeybase{name: name, dir: dir, kb: kb}, nil
}

func (ekb EagerKeybase) CreateAccount(
	name, mnemonic, bip39Passwd, encryptPasswd string,
	account uint32, index uint32,
) (keys.Info, error) {
	return ekb.kb.CreateAccount(name, mnemonic, bip39Passwd, encryptPasswd, account, index)
}

func (ekb EagerKeybase) GetByAddress(address string) (keys.Info, error) {
	return ekb.kb.GetByNameOrAddress(address)
}

func (ekb EagerKeybase) GetByName(name string) (keys.Info, error) {
	return ekb.kb.GetByName(name)
}

func (ekb EagerKeybase) Sign(name, passphrase string, msg []byte) ([]byte, crypto.PubKey, error) {
	return ekb.kb.Sign(name, passphrase, msg)
}

func (ekb EagerKeybase) Close() {
	ekb.kb.CloseDB()
}

func (ekb EagerKeybase) List() ([]keys.Info, error) {
	return ekb.kb.List()
}

type Txbase struct {
	name string
	dir  string
	db   tm2db.DB
}

func NewTxbase(rootDir, name string) (Txbase, error) {

	dir := filepath.Join(rootDir, "data")

	if err := os.EnsureDir(dir, 0o700); err != nil {
		panic(fmt.Sprintf("failed to create Keybase directory: %s", err))
	}
	db, err := tm2db.NewDB(name, dbBackend, dir)
	if err != nil {
		return Txbase{}, err
	}

	return Txbase{name: name, dir: dir, db: db}, nil
}

// the key is signer account address
func (tb Txbase) SetTxs(addr crypto.Address, txs []std.Tx) error {

	mtxs,err := amino.MarshalJSON(txs)
	  if err != nil{

			panic("")
		}
	tb.db.SetSync([]byte(addr.String()), mtxs)

	return nil
}
func (tb Txbase) GetTxsByAddrS(addr string )([]std.Tx, error){
	bzTx := tb.db.Get([]byte(addr))
	if len(bzTx) == 0 {
		return nil, fmt.Errorf("key with address %s not found", addr)
	}
	var txs []std.Tx

	err := amino.UnmarshalJSON(bzTx, &txs)
	if err != nil {
		return nil, err
	}
	return txs, nil

}
func (tb Txbase) GetTxs(addr crypto.Address) ([]std.Tx, error) {
    return  tb.GetTxsByAddrS(addr.String())
}
func (tb Txbase) Delete(addr crypto.Address){

	tb.db.DeleteSync([]byte(addr.String()))

}

type TxInfo struct{

	addr string
	txs []std.Tx

}

func (tb Txbase) List() []TxInfo {
	var res []TxInfo
	iter := tb.db.Iterator(nil, nil)
	defer iter.Close()
	for ; iter.Valid(); iter.Next() {
		key := string(iter.Key())

			txs, err := tb.GetTxsByAddrS(key)
			if err != nil {
				continue
			}
			info := TxInfo{
				addr: key,
				txs: txs,
			}
			res = append(res, info)

	}
	return res
}

func (tb Txbase) Close() {
	tb.db.Close()
}
