package p_crossrealm

import "std"

type Stringer interface {
	String() string
}

type Container struct {
	A int
	B Stringer
}

func (c *Container) Touch() *Container {
	c.A += 1
	return c
}

func (c *Container) Print() {
	println("A:", c.A)
	if c.B == nil {
		println("B: undefined")
	} else {
		println("B:", c.B.String())
	}
}

func CurrentRealm() std.Realm {
	return std.CurrentRealm()
}
