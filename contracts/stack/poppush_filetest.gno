package main

import (
	"strconv"
	"strings"
)

func main() {
	Pop()
	Push()
	println(Render(""))

}

var slice = []string{"undead-element"}

func Pop() {
	slice = slice[:len(slice)-1]
}

func Push() {
	slice = append(slice, "new-element")
}

func Render(path string) string {
	str := strings.Join(slice, ", ") + "\n\n"
	str += "len: " + strconv.Itoa(len(slice)) + "\n\n"
	return str
}

// Output:
// new-element
//
// len: 1
