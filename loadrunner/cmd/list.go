package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/gnolang/gno/tm2/pkg/commands"
)

type listCfg struct{}

func (c *listCfg) RegisterFlags(fs *flag.FlagSet) {}

func newListCmd() *commands.Command {
	cfg := &listCfg{}
	return commands.NewCommand(
		commands.Metadata{
			Name:       "list",
			ShortUsage: "list [signer address]",
			ShortHelp:  "list yet to broadcast txs stored locally",
		},
		cfg,
		func(_ context.Context, args []string) error {
			return list(args)
		},
	)
}

func list(args []string) error {
	// Txbase store signed Txs before it is broadcasted.
	txbase, err := NewTxbase(RootDir, defaultTxDBName)
	if err != nil {
		return err
	}
	defer txbase.Close()

	if len(args) == 0 {
		txinfos := txbase.List()
		fmt.Printf("%d entries in txs.db \n", len(txinfos))
		return nil
	}
	addr := args[0]

	txs, err := txbase.GetTxsByAddrS(addr)
	if err != nil {
		return err
	}

	fmt.Printf("Address \t\t\t\t\tGas wanted \tFee\n")
	for _, tx := range txs {
		fmt.Printf("%s \t%d \t%s\n", addr, tx.Fee.GasWanted, tx.Fee.GasFee)
	}

	return nil
}
