package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/gnolang/gno/gno.land/pkg/gnoland"
	vmm "github.com/gnolang/gno/gno.land/pkg/sdk/vm"
	gno "github.com/gnolang/gno/gnovm/pkg/gnolang"
	"github.com/gnolang/gno/gnovm/pkg/gnomod"
	"github.com/gnolang/gno/telemetry"
	"github.com/gnolang/gno/tm2/pkg/amino"
	abci "github.com/gnolang/gno/tm2/pkg/bft/abci/types"
	"github.com/gnolang/gno/tm2/pkg/bft/node"
	"github.com/gnolang/gno/tm2/pkg/bft/privval"
	bft "github.com/gnolang/gno/tm2/pkg/bft/types"
	"github.com/gnolang/gno/tm2/pkg/commands"
	"github.com/gnolang/gno/tm2/pkg/crypto"
	osm "github.com/gnolang/gno/tm2/pkg/os"
	"github.com/gnolang/gno/tm2/pkg/sdk"
	"github.com/gnolang/gno/tm2/pkg/std"
	"github.com/gnolang/gno/tm2/pkg/store"
)

type startCfg struct {
	baseCfg
	skipFailingGenesisTxs bool
	skipStart             bool
	genesisBalancesFile   string
	genesisTxsFile        string
	chainID               string
	genesisRemote         string
	genesisMaxVMCycles    int64
	pruneStrategy         string
	persistentPeers       string
	config                string
}

func newStartCmd(bc baseCfg) *commands.Command {
	cfg := startCfg{
		baseCfg: bc,
	}

	return commands.NewCommand(
		commands.Metadata{
			Name:       "start",
			ShortUsage: "start [flags]",
			ShortHelp:  "Run the full node",
		},
		&cfg,
		func(_ context.Context, args []string) error {
			return execStart(cfg, args)
		},
	)
}

func (c *startCfg) RegisterFlags(fs *flag.FlagSet) {
	fs.BoolVar(
		&c.skipFailingGenesisTxs,
		"skip-failing-genesis-txs",
		false,
		"don't panic when replaying invalid genesis txs",
	)

	fs.BoolVar(
		&c.skipStart,
		"skip-start",
		false,
		"quit after initialization, don't start the node",
	)

	fs.StringVar(
		&c.genesisBalancesFile,
		"genesis-balances-file",
		"./genesis/genesis_balances.txt",
		"initial distribution file",
	)

	fs.StringVar(
		&c.genesisTxsFile,
		"genesis-txs-file",
		"./genesis/genesis_txs.txt",
		"initial txs to replay",
	)

	fs.StringVar(
		&c.chainID,
		"chainid",
		"dev",
		"the ID of the chain",
	)

	fs.StringVar(
		&c.genesisRemote,
		"genesis-remote",
		"localhost:26657",
		"replacement for '%%REMOTE%%' in genesis",
	)

	fs.Int64Var(
		&c.genesisMaxVMCycles,
		"genesis-max-vm-cycles",
		10_000_000,
		"set maximum allowed vm cycles per operation. Zero means no limit.",
	)
	fs.StringVar(
		&c.pruneStrategy,
		"prune",
		"syncable",
		"set the state prune strategy: nothing | everything | syncable",
	)

	fs.StringVar(
		&c.persistentPeers,
		"peers",
		"",
		"persistent peers: node_id@host:port,...",
	)
}

func initTelemetry(ctx context.Context) error {
	var options []telemetry.Option

	if os.Getenv("TELEM_METRICS_ENABLED") == "true" {
		options = append(options, telemetry.WithOptionMetricsEnabled())
	}
	if os.Getenv("TELEM_TRACES_ENABLED") == "true" {
		options = append(options, telemetry.WithOptionTracesEnabled())
	}
	if portString := os.Getenv("TELEM_PORT"); portString != "" {
		port, err := strconv.ParseUint(portString, 10, 64)
		if err != nil {
			return fmt.Errorf("invalid port: %w", err)
		}

		options = append(options, telemetry.WithOptionPort(port))
	}
	if os.Getenv("TELEM_USE_FAKE_METRICS") == "true" {
		options = append(options, telemetry.WithOptionFakeMetrics())
	}

	// The string options can be added by default. Their absence would yield the same result
	// as if the option were excluded altogether.
	options = append(options, telemetry.WithOptionMeterName(os.Getenv("TELEM_METER_NAME")))
	options = append(options, telemetry.WithOptionExporterEndpoint(os.Getenv("TELEM_EXPORTER_ENDPOINT")))
	options = append(options, telemetry.WithOptionServiceName(os.Getenv("TELEM_SERVICE_NAME")))

	return telemetry.Init(ctx, options...)
}

func execStart(c startCfg, args []string) error {
	rootDir := c.rootDir
	tmcfg := &c.baseCfg.tmConfig

  
  // Attempt to initialize telemetry. If the enviroment variables required to initialize
	// telemetry are not set, then the initialization will do nothing.
	ctx := context.Background()
	if err := initTelemetry(ctx); err != nil {
		return fmt.Errorf("error initializing telemetry: %w", err)
	}

	// create priv validator first.
	// need it to generate genesis.json
	newPrivValKey := tmcfg.PrivValidatorKeyFile()
	newPrivValState := tmcfg.PrivValidatorStateFile()
	priv := privval.LoadOrGenFilePV(newPrivValKey, newPrivValState)

	// write genesis file if missing.
	genesisFilePath := filepath.Join(rootDir, tmcfg.Genesis)
	if !osm.FileExists(genesisFilePath) {
		genDoc := makeGenesisDoc(
			priv.GetPubKey(),
			c.chainID,
			c.genesisBalancesFile,
			loadGenesisTxs(c.genesisTxsFile, c.chainID, c.genesisRemote),
		)
		writeGenesisFile(genDoc, genesisFilePath)
	}

	// create application and node.
	gnoApp, err := gnoland.NewApp(rootDir, c.skipFailingGenesisTxs, logger, c.genesisMaxVMCycles)
	if err != nil {
		return fmt.Errorf("error in creating new app: %w", err)
	}

	// prune nothing is the archive node setting
	// prune syncable is the default node setting. It keeps the lastest 100 transactions and everything 1000th tx
	prune := store.NewPruningOptionsFromString(c.pruneStrategy)
	pruningOpt := sdk.SetPruningOptions(prune)
	gnoBaseApp := gnoApp.(*sdk.BaseApp)
	pruningOpt(gnoBaseApp)

	tmcfg.LocalApp = gnoBaseApp
	tmcfg.P2P.PersistentPeers = c.persistentPeers

	gnoNode, err := node.DefaultNewNode(tmcfg, logger)
	if err != nil {
		return fmt.Errorf("error in creating node: %w", err)
	}

	fmt.Println("Node created.")

	if c.skipStart {
		fmt.Println("'--skip-start' is set. Exiting.")

		return nil
	}

	if err := gnoNode.Start(); err != nil {
		return fmt.Errorf("error in start node: %w", err)
	}

	// run forever
	osm.TrapSignal(func() {
		if gnoNode.IsRunning() {
			_ = gnoNode.Stop()
		}
	})

	select {} // run forever
}

// Makes a local test genesis doc with local privValidator.
func makeGenesisDoc(
	pvPub crypto.PubKey,
	chainID string,
	genesisBalancesFile string,
	genesisTxs []std.Tx,
) *bft.GenesisDoc {
	gen := &bft.GenesisDoc{}

	gen.GenesisTime = time.Now()
	gen.ChainID = chainID
	gen.ConsensusParams = abci.ConsensusParams{
		Block: &abci.BlockParams{
			// TODO: update limits.
			MaxTxBytes:   1000000,  // 1MB,
			MaxDataBytes: 2000000,  // 2MB,
			MaxGas:       10000000, // 10M gas
			TimeIotaMS:   100,      // 100ms
		},
	}
	gen.Validators = []bft.GenesisValidator{
		{
			Address: pvPub.Address(),
			PubKey:  pvPub,
			Power:   10,
			Name:    "testvalidator",
		},
	}

	// Load distribution.
	balances := loadGenesisBalances(genesisBalancesFile)
	// debug: for _, balance := range balances { fmt.Println(balance) }

	// Load initial packages from examples.
	test1 := crypto.MustAddressFromString("g1jg8mtutu9khhfwc4nxmuhcpftf0pajdhfvsqf5")
	txs := []std.Tx{}

	// List initial packages to load from examples.
	pkgs, err := gnomod.ListPkgs(filepath.Join("..", "examples"))
	if err != nil {
		panic(fmt.Errorf("listing gno packages: %w", err))
	}

	// Sort packages by dependencies.
	sortedPkgs, err := pkgs.Sort()
	if err != nil {
		panic(fmt.Errorf("sorting packages: %w", err))
	}

	// Filter out draft packages.
	nonDraftPkgs := sortedPkgs.GetNonDraftPkgs()

	for _, pkg := range nonDraftPkgs {
		// open files in directory as MemPackage.
		memPkg := gno.ReadMemPackage(pkg.Dir, pkg.Name)

		var tx std.Tx
		tx.Msgs = []std.Msg{
			vmm.MsgAddPackage{
				Creator: test1,
				Package: memPkg,
				Deposit: nil,
			},
		}
		tx.Fee = std.NewFee(50000, std.MustParseCoin("1000000ugnot"))
		tx.Signatures = make([]std.Signature, len(tx.GetSigners()))
		txs = append(txs, tx)
	}

	// load genesis txs from file.
	txs = append(txs, genesisTxs...)

	// construct genesis AppState.
	gen.AppState = gnoland.GnoGenesisState{
		Balances: balances,
		Txs:      txs,
	}
	return gen
}

func writeGenesisFile(gen *bft.GenesisDoc, filePath string) {
	err := gen.SaveAs(filePath)
	if err != nil {
		panic(err)
	}
}

func loadGenesisTxs(
	path string,
	chainID string,
	genesisRemote string,
) []std.Tx {
	txs := []std.Tx{}
	txsBz := osm.MustReadFile(path)
	txsLines := strings.Split(string(txsBz), "\n")
	for _, txLine := range txsLines {
		if txLine == "" {
			continue // skip empty line
		}

		// patch the TX
		txLine = strings.ReplaceAll(txLine, "%%CHAINID%%", chainID)
		txLine = strings.ReplaceAll(txLine, "%%REMOTE%%", genesisRemote)

		var tx std.Tx
		amino.MustUnmarshalJSON([]byte(txLine), &tx)
		txs = append(txs, tx)
	}

	return txs
}

func loadGenesisBalances(path string) []string {
	// each balance is in the form: g1xxxxxxxxxxxxxxxx=100000ugnot
	balances := []string{}
	content := osm.MustReadFile(path)
	lines := strings.Split(string(content), "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)

		// remove comments.
		line = strings.Split(line, "#")[0]
		line = strings.TrimSpace(line)

		// skip empty lines.
		if line == "" {
			continue
		}

		parts := strings.Split(line, "=")
		if len(parts) != 2 {
			panic("invalid genesis_balance line: " + line)
		}

		balances = append(balances, line)
	}
	return balances
}
