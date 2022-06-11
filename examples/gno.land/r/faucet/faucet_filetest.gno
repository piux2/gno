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
		test3addr  = testutils.TestAddress("test3")
	)
	banker := std.GetBanker(std.BankerTypeReadonly)
	std.TestIssueCoins(faucetaddr, std.Coins{{"gnot", 1000}})
	std.TestSetOrigPkgAddr(faucetaddr)

	showBalances := func() string {
		mainbal := banker.GetCoins(mainaddr)
		faucetbal := banker.GetCoins(faucetaddr)
		test1bal := banker.GetCoins(test1addr)
		test2bal := banker.GetCoins(test2addr)
		test3bal := banker.GetCoins(test3addr)
		return fmt.Sprintf("main=%q, faucet=%q, test1=%q, test2=%q test3=%q",
			mainbal.String(), faucetbal.String(), test1bal.String(), test2bal.String(), test3bal.String())
	}

	println("before:", showBalances())

	// TODO: Approve an address to send
	// TODO: use that address to reward another one

	// simulate a Deposit call.
	std.TestSetOrigSend(std.Coins{{"gnot", 300}}, nil)
	faucet.Reward(test1addr)
	faucet.Reward(test1addr)
	faucet.Reward(test2addr)

	println("after:", showBalances())

	// simulate a Render().
	ret := faucet.Render("")
	println("---")
	println("Render: ", ret)
	println("---")
	println("done")
}

// Output:
// before: main="200gnot", faucet="1000gnot", test1="", test2="" test3=""
// undefined
// undefined
// undefined
// after: main="200gnot", faucet="700gnot", test1="200gnot", test2="100gnot" test3=""
// ---
// Render:  This faucet is managed by the community.
//
// To get some tokens, you should find someone allowed to use it to send you some tokens.
//
// Balance: 700gnot
// Rewarded: 300gnot (in 3 times)
//
// ---
// done

// Realm:
