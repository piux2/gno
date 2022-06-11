package faucet

import (
	"std"

	"gno.land/r/users"
)

// Sets `amount` as the allowance of `spenderName`.
func Approve(spenderName string, amount uint64) {
	std.AssertOriginCall()
	caller := std.GetOrigCaller()

	// TODO: if admin

	if amount < 0 {
		panic("invalid amount")
	}

	spender := users.GetUserByName(spenderName)
	if spender == nil {
		panic("unknown spender")
	}

	spenderaddr := spender.Address().String()
	agents.Set(spenderaddr, amount)
}
