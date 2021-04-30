package main

import "fmt"

func main() {
	ch := make(chan int, 10)
	go fib2(ch, cap(ch))
	for i := range ch {
		fmt.Println(i)
	}
}

func fib2(ch chan int, n int) {
	x, y := 1, 1
	for i := 0; i < n; i++ {
		ch <- x
		x, y = y, x+y
	}
	close(ch)
}
