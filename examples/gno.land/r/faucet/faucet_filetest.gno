package main

import (
	"fmt"
	"std"

	"gno.land/p/testutils"
	"gno.land/r/faucet"
)

func main() {
	var (
		mainaddr   = std.TestDerivePkgAddr("main")
		faucetaddr = std.TestDerivePkgAddr("gno.land/r/faucet")
		test1addr  = testutils.TestAddress("test1")
		test2addr  = testutils.TestAddress("test2")
	)
	banker := std.GetBanker(std.BankerTypeReadonly)
	std.TestIssueCoins(faucetaddr, std.Coins{{"gnot", 1000}})
	std.TestSetOrigPkgAddr(faucetaddr)

	showBalances := func() string {
		mainbal := banker.GetCoins(mainaddr)
		faucetbal := banker.GetCoins(faucetaddr)
		test1bal := banker.GetCoins(test1addr)
		test2bal := banker.GetCoins(test2addr)
		return fmt.Sprintf("main=%q, faucet=%q, test1=%q, test2=%q",
			mainbal.String(), faucetbal.String(), test1bal.String(), test2bal.String())
	}

	println("before:", showBalances())

	// simulate a Deposit call.
	std.TestSetOrigSend(std.Coins{{"gnot", 100}}, nil)
	faucet.Reward(test1addr)

	println("after:", showBalances())

	// simulate a Render().
	ret := faucet.Render("")
	println(ret)
	println("done")
}

func printBalances() {

}

// Output:
// before: main="200gnot", faucet="1000gnot", test1="", test2=""
// after: main="200gnot", faucet="900gnot", test1="100gnot", test2=""
//
// This faucet is managed by the community.
// To get some tokens, you should find someone allowed to use it to send you some tokens.
//
// done

// Realm:
