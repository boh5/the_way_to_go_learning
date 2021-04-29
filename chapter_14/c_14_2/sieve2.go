package main

import "fmt"

func main() {
	primes := Sieve()
	for {
		fmt.Println(<-primes)
	}
}

func Generate() chan int {
	ch := make(chan int)
	go func() {
		for i := 2; ; i++ {
			ch <- i
		}
	}()
	return ch
}

func Filter(in chan int, prime int) chan int {
	out := make(chan int)
	go func() {
		for {
			if i := <-in; i%prime != 0 {
				out <- i
			}
		}
	}()
	return out
}

func Sieve() chan int {
	out := make(chan int)
	go func() {
		ch := Generate()
		for {
			prime := <-ch
			ch = Filter(ch, prime)
			out <- prime
		}
	}()
	return out
}
