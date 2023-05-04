package main

type nat []Word
type Word uint

func (n nat) add() bool {
	return true
}

func main() {
	var abs nat
	abs = []Word{0}
	println(abs.add())
}

// Output:
// true
