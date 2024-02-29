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
	"github.com/gnolang/gno/tm2/pkg/crypto"
	"github.com/gnolang/gno/tm2/pkg/crypto/keys"

	"github.com/gnolang/gno/gno.land/pkg/sdk/vm"
	"github.com/gnolang/gno/tm2/pkg/bft/rpc/client"
	"github.com/gnolang/gno/tm2/pkg/std"
)

const (
	ChainID = "dev"
	numMsgs = 1    // number messages per tx
)

type signCfg struct{}

func (c *signCfg) RegisterFlags(fs *flag.FlagSet) {}

func newSignCmd() *commands.Command {
	cfg := &signCfg{}
	return commands.NewCommand(
		commands.Metadata{
			Name:       "sign",
			ShortUsage: "sign [first n-th signers]",
			ShortHelp:  "sign txs and store locally",
		},
		cfg,
		func(_ context.Context, args []string) error {
			return sign(args[0])
		},
	)
}

type SignTxTask struct {
	SignerKey      string
	Account         std.BaseAccount
	encryptPassword string
	keybase         *EagerKeybase
	txs             []std.Tx
	txDB            *Txbase
}

func (s SignTxTask) Key() string {
	return s.SignerKey
}

func (s SignTxTask) Execute() error {
	info, err := s.keybase.GetByName(s.SignerKey)
	if err != nil {
		return err
	}

	signedTxs, err := signTx(info, s.Account, s.txs, s.keybase)
	if err != nil {
		return err
	}
	fmt.Println("worker", s.SignerKey)
	(*s.txDB).SetTxs(info.GetAddress(), signedTxs)
	return nil
}

func sign(num string) error {
	var numWorkers int = 1000

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

	// we will assign each key account a task to sign the transactions.
	// n is number of key accounts.
	n, err := strconv.Atoi(num)
	if err != nil {
		return nil
	}
	// check and update each account sequence	against the node
	signTasks, err := prepareSignTasks(kb,txbase, n)
	if err != nil {
		return err
	}

	if l :=len(signTasks) ; n > l {
		n  = l
	}


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
	for _, t := range signTasks{


		tasks <- t
	}

	wg.Wait()
	close(tasks)
	close(results)

	println()
	fmt.Printf("signed %d message in %s/data/txs.db\n", n*numMsgs, RootDir)
	return nil
}

func newTxs() []std.Tx {
	// TODO: max gas, max tx size and max msg size
	pkgPath := "gno.land/r/x/benchmark/load"
	fn := "AddPost"
	args := []string{"hello","world"}
/*
	args := []string{
		"Weather Outlook: Nov 1 - Nov 7, 2024: A Week of Changing Skies", "Today's comprehensive weather forecast promises a dynamic and engaging experience for all, blending a mix of atmospheric conditions that cater to a wide array of preferences and activities. As dawn breaks, residents can anticipate a refreshing and crisp morning with temperatures gently rising from a cool 55°F, creating an invigorating start to the day. The early hours will see a soft, dew-kissed breeze whispering through the streets, carrying the fresh scent of blooming flowers and newly cut grass, setting a serene tone for the day ahead.\n\nBy mid-morning, the sun, in its splendid glory, will begin to assert its presence, gradually elevating temperatures to a comfortable 75°F. The skies, adorned with a few scattered clouds, will paint a picturesque backdrop, ideal for outdoor enthusiasts eager to embrace the day's warmth. Whether it's a leisurely stroll in the park, an adventurous hike through nearby trails, or simply enjoying a quiet moment in the sun, the conditions will be perfectly aligned for an array of outdoor pursuits.\n\nAs the day progresses towards noon, expect the gentle morning breeze to evolve into a more pronounced wind, adding a refreshing counterbalance to the midday sun's warmth. This perfect harmony between the breeze and sunlight offers an optimal environment for sailing and kite-flying, providing just the right amount of lift and drift for an exhilarating experience.\n\nThe afternoon promises a continuation of the day's pleasant conditions, with the sun reigning supreme and the temperature peaking at a delightful 80°F. It's an ideal time for community sports, gardening, or perhaps an outdoor picnic, allowing friends and family to gather and make the most of the splendid weather.\n\nHowever, as we transition into the evening, anticipate a slight shift in the atmosphere. The temperature will gently dip, creating a cool and comfortable setting, perfect for al fresco dining or a serene walk under the starlit sky. The night will conclude with a mild 60°F, ensuring a peaceful and restful end to a day filled with diverse weather experiences.\n\nIn summary, today's weather forecast offers something for everyone, from the early risers seeking tranquility in the morning's embrace to the night owls looking to unwind under the cool evening air. It's a day to revel in the outdoors, pursue a myriad of activities, and simply enjoy the natural beauty that surrounds us.",
	}
*/
	msgs := []std.Msg{}
	gaswanted := int64(10000000)
	gasfee := std.Coin{
		Denom:  "ugnot",
		Amount: 1,
	}

	msg := vm.MsgCall{
		Caller:  crypto.Address{},
		PkgPath: pkgPath,
		Func:    fn,
		Args:    args,
	}

	for i := 0; i < numMsgs; i++ {
		msgs = append(msgs, msg)
	}

	tx := std.Tx{
		Msgs:       msgs,
		Fee:        std.NewFee(gaswanted, gasfee),
		Signatures: nil,
	}
	return []std.Tx{tx}
}

func signTx(signer keys.Info, account std.BaseAccount, txs []std.Tx, kb *EagerKeybase) ([]std.Tx, error) {
	// sign tx
	accountNumber := account.AccountNumber
	sequence := account.Sequence
	signerAddress := signer.GetAddress()
	var signedTxs []std.Tx

	for _, tx := range txs {
		// fill the msg caller
		for i, msg := range tx.Msgs {
			switch msg2 := msg.(type) {
			case vm.MsgCall:
				msg2.Caller = signerAddress
				tx.Msgs[i] = msg2
			default:
				return nil, fmt.Errorf("Invalid msg type %+v", msg)
			}
		}

		signers := tx.GetSigners()
		// derive sign doc bytes.
		signbz := tx.GetSignBytes(ChainID, accountNumber, sequence)

		sig, pub, err := kb.Sign(signerAddress.String(), encryptPassword, signbz)
		if err != nil {
			return nil, err
		}
		addr := pub.Address()
		found := false

		for i := range signers {
			if signers[i] == addr {
				found = true
				tx.Signatures = append(tx.Signatures, std.Signature{
					PubKey:    pub,
					Signature: sig,
				})
			}
		}

		if !found {
			err := fmt.Errorf("addr %v not in signer set %v\n", addr, signers)
			return nil, err

		}
		signedTxs = append(signedTxs, tx)

		accountNumber := account.AccountNumber
		sequence := account.Sequence
		signerAddress := signer.GetAddress()
		fmt.Printf("sign: %s, account: %d, sequence: %d\n", signerAddress, accountNumber, sequence)
		sequence++
	}
	return signedTxs, nil
}

// update sequence number of first n-th accounts
func prepareSignTasks(kb EagerKeybase,txbase Txbase, n int) ([]Task, error) {
	infos, err := kb.List()
	if err != nil {
		return nil, err
	}
	// max is len(infos)
	l := len(infos)
	if n > l {
		return nil, fmt.Errorf("There are only %d accounts, %d is too big", l, n)
	}
	var tasks []Task

	// for _, info := range infos {
	for i := 0; i < n; i++ {
		// update account's information
		key := SignerKeyPrefix + strconv.Itoa(i)
		info,err := kb.GetByName(key)
		if err != nil{
			continue
		}
		qopts := &queryOption{
			path: fmt.Sprintf("auth/accounts/%s", info.GetAddress()),
		}

		qres, err := queryHandler(qopts)
		if err != nil {
			return nil, fmt.Errorf("query account %w\n", err)
		}
		var qret struct{ BaseAccount std.BaseAccount }
		err = amino.UnmarshalJSON(qres.Response.Data, &qret)
		if err != nil {
			return nil, err
		}
		fmt.Printf("update account: %v\n", qret.BaseAccount)

		// create  Tasks.
		txs := newTxs()
		t := SignTxTask{
			SignerKey:      key,
			Account:         qret.BaseAccount,
			encryptPassword: encryptPassword,
			txs:             txs,
			keybase:         &kb,
			txDB:            &txbase,
		}
		tasks = append(tasks, t)

	}




	return tasks, nil
}

type queryOption struct {
	data   string
	height int64
	prove  bool
	// internal
	path string
}

const remote = "127.0.0.1:26657"

func queryHandler(qopt *queryOption) (*ctypes.ResultABCIQuery, error) {
	data := []byte(qopt.data)
	opts2 := client.ABCIQueryOptions{
		// Height: height, XXX
		// Prove: false, XXX
	}
	cli := client.NewHTTP(remote, "/websocket")
	qres, err := cli.ABCIQueryWithOptions(
		qopt.path, data, opts2)
	if err != nil {
		return nil, fmt.Errorf("querying %w\n", err)
	}

	return qres, nil
}
