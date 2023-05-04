package main

type nat []Word
type Word uint

func main() {
	var abs nat
	abs = []Word{1, Word(2)}
	println(abs)
}

// Output:
// slice[(1 main.Word),(2 main.Word)]
