package main


type nat []int
type nat2 []int

func main() {

	var nlist = nat{0}
	var n2list = nat2{1}

	copy(nlist, n2list)


	println(nlist)


}

// Output:
// (slice[(1 int)] main.nat)
