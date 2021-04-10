package main

import "fmt"

func main() {
	forLoop()
	forLoopWithoutFor()
}
func forLoop() {
	for i := 0; i < 15; i++ {
		fmt.Println(i)
	}
}

func forLoopWithoutFor() {
	i := 0
START:
	fmt.Println(i)
	i++
	if i < 15 {
		goto START
	}
}
