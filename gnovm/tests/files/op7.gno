package main

type T int

func (t T) Error() string { return "T: error" }

var invalidT T

func main() {
	var err error
	if err > invalidT {
		println("ok")
	}
}

// Error:
// main/files/op7.gno:11:5: operator > not defined on: InterfaceKind

// TypeCheckError:
// main/files/op7.gno:11:5: invalid operation: err > invalidT (operator > not defined on interface)
