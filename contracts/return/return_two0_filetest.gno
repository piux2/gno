package main

type nat []int

func x() ([]int, nat) {
	a := nat{0}
	b := []int{1}
	return a, b
}

func main() {
	a, b := x()
	println(a)
	println(b)

}

// Output:
// slice[(0 int)]
// (slice[(1 int)] main.nat)
