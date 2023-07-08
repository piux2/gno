package main

type nat []int

func x() (interface{}, interface{}) {
	a := nat{0}

	return a, a
}

func main() {
	var a, b nat

	a, b = x() // This should fail

	println(a)
	println(b)

}

// Error:
// main/contracts/regression/return_two1e_filetest.gno:14: cannot use interface{} as []int
