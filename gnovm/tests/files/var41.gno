// Genuinely cyclic var declarations must still be rejected.
package main

var (
	A = B
	B = A
)

func main() {
	println(A, B)
}

// Error:
// main/var41.gno:4:2-7: invalid recursive value: A -> B -> A
