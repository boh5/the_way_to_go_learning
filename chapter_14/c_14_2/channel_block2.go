package main

import (
	"fmt"
	"time"
)

func main() {
	stream := Pump()
	go Suck(stream)
	time.Sleep(1e9)
}

func Pump() chan int {
	ch := make(chan int)
	go func() {
		for i := 0; ; i++ {
			ch <- i
		}
	}()
	return ch
}

func Suck(ch chan int) {
	for {
		fmt.Println(<-ch)
	}
}
