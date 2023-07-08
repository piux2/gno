package main

func main() {
	var n1 []int
	// transcribe n1 = []int{0} to tmp:= []int{0}, n1=temp
	// n1 = []int{0}
	tmp := []int{0}
	n1 = tmp

	println(n1)
}

// Output:
// slice[(0 int)]
