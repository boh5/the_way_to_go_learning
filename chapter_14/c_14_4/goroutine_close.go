package main

import "fmt"

func main() {
	ch := make(chan int)
	go tel(ch, 100)
	for i := range ch {
		fmt.Println(i)
	}
}

func tel(ch chan int, max int) {
	for i := 0; i < max; i++ {
		ch <- i
	}
	close(ch)
}
