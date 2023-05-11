package main

type nat []Word
type Word uint

func (n nat) add(ws []Word) nat {

	return ws
}

func (n nat) add2(nt nat) []Word {

	return nt
}

func (n nat) add3(ws []Word) interface{} {

	return ws
}

func (n nat) add4(nt nat) interface{} {

	return nt
}

func main() {

	var abs nat
	abs = []Word{0}

	println(abs)
	println(abs.add(abs))
	println(abs.add2(abs))
	println(abs.add3(abs))
	println(abs.add4(abs))
	println()

	abs = nat{1}

	println(abs)
	println(abs.add(abs))
	println(abs.add2(abs))
	println(abs.add3(abs))
	println(abs.add4(abs))
	println()

	def := []Word{2}
	println(def)
	println(abs.add(def))
	println(abs.add2(def))
	println(abs.add3(def))
	println(abs.add4(def))
	println()

	def = nat{3}
	println(def)
	println(abs.add(def))
	println(abs.add2(def))
	println(abs.add3(def))
	println(abs.add4(def))
	println()

}

// Output:
// (slice[(0 main.Word)] main.nat)
// (slice[(0 main.Word)] main.nat)
// slice[(0 main.Word)]
// slice[(0 main.Word)]
// (slice[(0 main.Word)] main.nat)
//
// (slice[(1 main.Word)] main.nat)
// (slice[(1 main.Word)] main.nat)
// slice[(1 main.Word)]
// slice[(1 main.Word)]
// (slice[(1 main.Word)] main.nat)
//
// slice[(2 main.Word)]
// (slice[(2 main.Word)] main.nat)
// slice[(2 main.Word)]
// slice[(2 main.Word)]
// (slice[(2 main.Word)] main.nat)
//
// slice[(3 main.Word)]
// (slice[(3 main.Word)] main.nat)
// slice[(3 main.Word)]
// slice[(3 main.Word)]
// (slice[(3 main.Word)] main.nat)
