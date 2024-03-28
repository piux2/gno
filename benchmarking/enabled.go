package benchmarking

import (
	"fmt"
	"os"
	"strconv"

)

const (
	KEEPER_CALL   = "call"
	KEEPER_INIT   = "init"
	KEEPER_ADDPKG = "addpkg"
)

type  filter int64

const (
	None filter = iota
	CPU
	Store
)



var gFilter filter

// There are two control points to isolate benchmarking scope.
// - Keeper entry points at Init, Msg_Call, Msg_AddPkg
var (
	Entry string
	// We set start cpu benchmarking for true after an OpCode executed
	StartCPU bool

	//We set start store benchmarking for true after an OpCode executed
	StartStore bool
	// we only turn OpCodeDetails on to understand the OpCode in benchmarking call flow. We turn it off for accurate measurement timing
	OpCodeDetails bool
)

var enabled bool

func Enabled() bool {
	return enabled
}

func Init(filepath string) {
	enabled = true
	initExporter(filepath)
	initStack()
	if os.Getenv("OPCODE_DETAILS") == "true" {
		OpCodeDetails = true
	}
	if os.Getenv("BENCHMARK_FILTER") != "" {
		filterType, err := strconv.ParseInt(os.Getenv("BENCHMARK_FILTER"), 10, 64)
		if err != nil {
			panic(fmt.Errorf("invalid benchmark filter: %w", err))
		}
		gFilter = filter(filterType)
	}

}

func IsCPU() bool {
	return gFilter == None || gFilter == CPU
}

func IsStore() bool {
	return gFilter == None || gFilter == Store
}
