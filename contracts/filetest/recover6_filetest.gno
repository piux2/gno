package main

import (
	"errors"
)

func main() {
	println(f(false))
	println(f(true))
}

func f(dopanic bool) (err error) {
	defer func() {
		if x := recover(); x != nil {
			println("recover start")
			println(x)
			err = x.(error)
			println(err)
			println("recover end")
		}
	}()
	q(dopanic)
	return
}

func q(dopanic bool) {
	if dopanic {
		panic(errors.New("wtf"))
	}
}

// Output:
// undefined
// recover start
// wtf
// wtf
// recover end
// wtf
