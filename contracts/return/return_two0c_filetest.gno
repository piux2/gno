package main

type nat []int

func x() (nat, []int) {
	a := nat{0}
	b := []int{1}
	return a, b
}

func main() {
	var a nat
	var b []int
	a, b = x()
	println(a)
	println(b)

}

// Output:
// (slice[(0 int)] main.nat)
// slice[(1 int)]
