package main

import "fmt"

func main() {
	var a [5]int
	b := a
	b[0] = 100
	fmt.Println(a, b)
}
