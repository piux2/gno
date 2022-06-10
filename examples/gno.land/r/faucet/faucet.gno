package faucet

import (
	"std"

	"gno.land/r/users"
)

func Reward(to std.Address) {
	//std.AssertOriginCall()
	//caller := std.GetOrigCaller()

	// if allowance

	send := std.Coins{{"gnot", 100}}
	banker := std.GetBanker(std.BankerTypeOrigSend)
	pkgaddr := std.GetOrigPkgAddr()
	//banker.SendCoins(pkgaddr, to, send)
}

func Render(path string) string {
	return "TODO"
}
