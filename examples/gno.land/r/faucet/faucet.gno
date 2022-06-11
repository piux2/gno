package faucet

import (
	"std"

	"gno.land/p/ufmt"
	"gno.land/r/users"
)

var (
	totalRewarded = std.Coin{"gnot", 0}
	uniqueRewards = 0
)

func Reward(to std.Address) {
	std.AssertOriginCall()
	caller := std.GetOrigCaller()

	amount := int64(100)

	// TODO: if allowance

	// TODO: log history

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
