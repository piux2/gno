// Multi-value decl where one split name feeds back into the same decl.
// A gets C, but C depends on A — genuine cycle, not a false positive.
package main

var (
	A, B = C, D
	C    = A
	D    = 2
)

func main() {
	println(A, B, C)
}

// Error:
// main/var45.gno:4:1-9:2: invalid recursive value: A -> C -> A
