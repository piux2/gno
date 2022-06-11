package faucet

import (
	"std"

	"gno.land/r/users"
)

func Reward(to std.Address) {
	std.AssertOriginCall()
	caller := std.GetOrigCaller()

	// TODO: if allowance

	// TODO: log history

	send := std.Coins{{"gnot", 100}}
	banker := std.GetBanker(std.BankerTypeOrigSend)
	pkgaddr := std.GetOrigPkgAddr()
	banker.SendCoins(pkgaddr, to, send)
}

func Render(path string) string {
	output := `
This faucet is managed by the community.
To get some tokens, you should find someone allowed to use it to send you some tokens.
`
	return output
}
