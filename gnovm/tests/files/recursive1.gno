package main

type S struct {
	T S
}

func main() {
	var a, b S
	println(a == b)
}

// Error:
// main/files/recursive1.gno:1:1: invalid recursive type: S -> S
