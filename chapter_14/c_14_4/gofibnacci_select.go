package main

import (
	"fmt"
	"os"
)

func main() {
	ch := make(chan int)
	stop := make(chan int)
	go fibS(ch, stop, 10)
	for {
		select {
		case v := <-ch:
			fmt.Println(v)
		case <-stop:
			os.Exit(0)
		}
	}
}

func fibS(ch, stop chan int, n int) {
	x, y := 1, 1
	for i := 0; i < n; i++ {
		ch <- x
		x, y = y, x+y
	}
	stop <- 1
}
