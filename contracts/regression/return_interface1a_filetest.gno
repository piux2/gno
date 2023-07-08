package main

// with the first parameter string. it is ok
// with either parameter is string, not ok
func x() (interface{}, string) {
	a := "0"
	b := "1"
	return a, b
}

func main() {
	var a string
	var b string
	a, b = x() // should fail
	println(a)
	println(b)

}

// Error:
// main/contracts/regression/return_interface1a_filetest.gno:14: cannot use interface{} as string
