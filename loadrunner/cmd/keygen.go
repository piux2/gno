// KeyGen
package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"sync"

	"github.com/gnolang/gno/tm2/pkg/commands"
)

const (
	SignerKeyPrefix = "loadrunner"
	encryptPassword = "loadrunner"
)

var RootDir string = "."

type keyGenCfg struct{}

func (c *keyGenCfg) RegisterFlags(fs *flag.FlagSet) {}

func newKeyGenCmd() *commands.Command {
	cfg := &keyGenCfg{}
	return commands.NewCommand(
		commands.Metadata{
			Name:       "keygen",
			ShortUsage: "keygen [num of keys]",
			ShortHelp:  "Generate n Keys",
		},
		cfg,
		func(_ context.Context, args []string) error {
			return keyGen(args)
		},
	)
}

type KeyGenTask struct {
	signerkey       string
	mnemonic        string
	bip39Passphrase string
	encryptPassword string
	account         uint32
	index           uint32
	keybase         *EagerKeybase
}

func (k KeyGenTask) Key() string {
	return k.signerkey
}

func (k KeyGenTask) Execute() error {
	_, err := (*k.keybase).CreateAccount(k.signerkey, k.mnemonic, k.bip39Passphrase, k.encryptPassword, k.account, k.index)
	if err != nil {
		return err
	}
	return nil
}

func keyGen(args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("%s\n", "please specify number of keys to generate.")
	}
	arg := args[0]
	var numWorkers int = 1000

	kb, err := NewEagerKeybase(RootDir)
	if err != nil {
		return err
	}
	defer kb.Close()

	// address g1rf78yealur05hf8gp7el435m8rtza8gkwe7lsx test2
	mnemonic := "pumpkin leaf essence pepper match ball will lens eyebrow fringe wheat place naive subject basket car carbon trigger pigeon quick melt garlic humble nephew"
	bip39Passphrase := ""

	var wg sync.WaitGroup
	n, err := strconv.Atoi(arg)
	if err != nil {
		return err
	}
	tasks := make(chan Task, n)
	results := make(chan Result, n) // Ensure the buffer is large enough to avoid blocking
	// start the worker routine

	// Start workers
	for w := 0; w < numWorkers; w++ {
		go worker(w, tasks, &wg, results)
		// go worker(w, tasks, &wg)
	}
	// start monitor
	go monitor(results, n)

	// jobs
	wg.Add(n)
	for i := 0; i < n; i++ {

		key := SignerKeyPrefix + strconv.Itoa(i)

		t := KeyGenTask{
			signerkey:       key,
			mnemonic:        mnemonic,
			bip39Passphrase: bip39Passphrase,
			encryptPassword: encryptPassword,
			account:         uint32(i),
			index:           uint32(0),
			keybase:         &kb,
		}

		tasks <- t
	}

	wg.Wait()
	close(tasks)
	close(results)

	println()
	fmt.Printf("generated %d keys in %s/data/keys.db", n, RootDir)

	genBalance(RootDir, kb)
	return nil
}

func genBalance(dir string, kb EagerKeybase) {
	output := "genesis_balances.txt"
	infos, err := kb.List()
	if err != nil {
		fmt.Println("can not list keybase")
	}

	p := filepath.Join(dir, output)
	file, err := os.Create(p)
	if err != nil {
		fmt.Println("can not create file:", output)
		return
	}
	defer file.Close()

	for _, info := range infos {

		addr := info.GetAddress().String()
		s := fmt.Sprintf("%s=1000000ugnot\n", addr)
		_, err := file.WriteString(s)
		if err != nil {
			fmt.Printf("Can not write %s, %v\n", s, err)
			return
		}
	}
}
