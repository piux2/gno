package main

type nat []int

func x() (nat, []int) {
	a := nat{0}
	b := []int{1}
	c := nat{2}
	return a, b
}

func main() {
	var n1, n2 nat
	var u1, u2 []int
	var a, b int

	n1, n2 = x()
	println(n1)
	println(n2)
	//println(n3)

	u1, u2 = x()
	println(u1)
	println(u2)

	n1, u2 = x()
	println(n1)
	println(u2)

	n2, u1 = x()

	println(n2)
	println(u1)

}

// Output:
// (slice[(0 int)] main.nat)
// (slice[(1 int)] main.nat)
// slice[(0 int)]
// slice[(1 int)]
// (slice[(0 int)] main.nat)
// slice[(1 int)]
// (slice[(0 int)] main.nat)
// slice[(1 int)]
