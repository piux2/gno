package main

type nat []int

func x() interface{} {
	a := nat{0}
	return a
}

func main() {
	var a nat

	a = x()

	println(a)

}

// Error:
// main/contracts/return/return_interface_filetest.gno:13: cannot use interface{} as []int
