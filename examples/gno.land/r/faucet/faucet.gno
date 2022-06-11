package faucet

import (
	"std"

	"gno.land/p/avl"
	"gno.land/p/ufmt"
)

var (
	totalRewarded = std.Coin{"gnot", 0}
	uniqueRewards = 0
	agents        = avl.NewMutTree() // std.Address -> allowance
)

func Reward(to std.Address) {
	std.AssertOriginCall()
	caller := std.GetOrigCaller()

	amount := int64(100)

	// TODO: log history?
	// TODO: limit per dest?

	agent, _ := agents.Get(caller.String())
	println(agent)
	// TODO: if allowance

	// TODO: allowance--

	send := std.Coins{{"gnot", amount}}
	totalRewarded.Amount += amount
	uniqueRewards++
	banker := std.GetBanker(std.BankerTypeOrigSend)
	pkgaddr := std.GetOrigPkgAddr()
	banker.SendCoins(pkgaddr, to, send)
}

func Render(path string) string {
	banker := std.GetBanker(std.BankerTypeOrigSend)
	balance := banker.GetCoins(std.GetOrigPkgAddr())

	output := "This faucet is managed by the community.\n\n"
	output += "To get some tokens, you should find someone allowed to use it to send you some tokens.\n\n"
	output += ufmt.Sprintf("Balance: %s\n", balance.String())
	output += ufmt.Sprintf("Rewarded: %s (in %d times)\n", totalRewarded.String(), uniqueRewards)
	return output
}
