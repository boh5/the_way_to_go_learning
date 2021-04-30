package main

import (
	"fmt"
	"os"
)

func main() {
	ch := make(chan int)
	stop := make(chan int)
	go Tel(ch, stop, 100)
	for {
		select {
		case v := <-ch:
			fmt.Println(v)
		case <-stop:
			os.Exit(0)
		}
	}
}

func Tel(ch, stop chan int, max int) {
	for i := 0; i < max; i++ {
		ch <- i
	}
	stop <- 1
}
