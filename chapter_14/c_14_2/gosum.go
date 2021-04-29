package main

import "fmt"

func main() {
	ch := make(chan int)
	go func(a, b int, ch chan int) {
		ch <- a + b
	}(1, 2, ch)

	fmt.Println(<-ch)
}
