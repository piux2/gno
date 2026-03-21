// Three-var cycle: A -> B -> C -> A must be rejected.
package main

var (
	A = B
	B = C
	C = A
)

func main() {
	println(A)
}

// Error:
// main/var43.gno:5:2-7: invalid recursive value: A -> B -> C -> A
