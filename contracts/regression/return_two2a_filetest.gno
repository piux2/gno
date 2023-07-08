package main

type nat []int

func x1() nat {
	a := nat{0}

	return a
}
func x2() []int {

	b := []int{1}
	return b
}

func main() {
	var n2 nat
	var u1 []int

	n2 = x2()

	println(n2)

	u1 = x1()
	println(u1)

}

// Output:
// (slice[(1 int)] main.nat)
// slice[(0 int)]
