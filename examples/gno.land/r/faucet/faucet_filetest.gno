package main

import (
	"std"

	"gno.land/r/faucet"
)

func main() {
	faucetAddr := std.TestDerivePkgAddr("gno.land/r/faucet")

	// print main balance before.
	mainaddr := std.TestDerivePkgAddr("main")
	banker := std.GetBanker(std.BankerTypeReadonly)
	mainbal := banker.GetCoins(mainaddr)
	println("main before:", mainbal) // plus OrigSend equals 300.

	// simulate a Deposit call.
	std.TestSetOrigPkgAddr(faucetAddr)
	std.TestIssueCoins(faucetAddr, std.Coins{{"gnot", 100}})
	std.TestSetOrigSend(std.Coins{{"gnot", 100}}, nil)
	res := faucet.Reward("1xblah")
	println("Deposit():", res)

	// print main balance after.
	mainbal = banker.GetCoins(mainaddr)
	println("main after:", mainbal) // still 300.

	// simulate a Render().
	ret := faucet.Render("")
	println(ret)
}

// Output:
// main before: 200gnot

// Realm:
