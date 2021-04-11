package main

import "fmt"

func main() {
	f1()
}

func f1() {
	fmt.Printf("In f1 at the top\n")
	defer f2()
	fmt.Printf("In f2 at the bottom!\n")
}

func f2() {
	fmt.Printf("f2: defferred until the end of the calling function!")
}
