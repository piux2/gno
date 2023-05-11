package main

import "fmt"

type nat []Word
type Word uint

func main() {
	var abs nat
	abs = []Word{0}
	println(abs)

}

// Output:
// (slice[(0 main.Word)] main.nat)
