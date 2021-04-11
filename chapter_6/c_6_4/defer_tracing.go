package main

import "fmt"

func main() {
	b()
}
func b() {
	trace("b")
	defer untrace("b")
	fmt.Println("in b")
	a()
}
func a() {
	trace("a")
	defer untrace("a")
	fmt.Println("in a")
}

func trace(s string) {
	fmt.Println("entering:", s)
}

func untrace(s string) {
	fmt.Println("leaving:", s)
}
