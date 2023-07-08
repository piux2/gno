package main

func x() (int, int) {

	return 1, 2
}

func main() {
	var a, b int
	a, _ = x()
	//t1, t2 := x()
	//a, b = t1, t2
	println(a)
	println(b)

}

// Output:
// 1
// 0
