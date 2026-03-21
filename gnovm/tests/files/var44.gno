// Multi-value decl where the two names swap — a genuine cycle.
package main

var (
	A, B = B, A
)

func main() {
	println(A, B)
}

// Error:
// main/var44.gno:4:1-6:2: invalid recursive value: A -> B -> A
