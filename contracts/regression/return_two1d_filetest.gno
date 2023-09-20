package main

type nat []int

func x() (interface{}, nat) {
	a := nat{0}
	b := []int{1}
	return a, b
}

func main() {
	var a nat
	var b []int
	a, b = x() // This should have failed.
	println(a)
	println(b)

}

// Output
