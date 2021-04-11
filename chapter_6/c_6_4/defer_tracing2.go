package main

import "fmt"

func main() {
	b2()
}

func b2() {
	defer un(trace2("b"))
	fmt.Println("in b")
	a2()
}

func a2() {
	defer un(trace2("a"))
	fmt.Println("in a")
}

func un(s string) {
	fmt.Println("leaving:", s)
}

func trace2(s string) string {
	fmt.Println("entering:", s)
	return s
}
