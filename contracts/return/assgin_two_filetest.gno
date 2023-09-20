package main

type nat []int
type nat2 []int

func main() {
	a := nat{0}
	b := nat2{1}
	a = b

	println(a)
	println(b)
}

// Error:
// main/contracts/return/assgin_two_filetest.gno:9: cannot use main.nat2 as main.nat without explicit conversion
