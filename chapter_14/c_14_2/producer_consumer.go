package main

import "fmt"

func main() {
	ch := make(chan int)
	done := make(chan bool)

	go producer(ch)
	go consumer(ch, done)
	<-done
}

func producer(ch chan<- int) {
	for i := 0; i < 91; i += 10 {
		ch <- i
	}
	close(ch)
}

func consumer(ch <-chan int, done chan<- bool) {
	for num := range ch {
		fmt.Println(num)
	}
	done <- true

}
