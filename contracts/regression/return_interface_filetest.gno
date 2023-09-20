package main

func x() (interface{}, interface{}) {
	a := 0
	return a, a
}

func main() {
	var a, b int
	a, b = x() // should have failed LINE 10
	println(a)
	println(b)

}

// Error:
// main/contracts/regression/return_interface_filetest.gno:10: cannot use interface{} as int
