package main

import "fmt"

func main() {
	for i := 0; i < 11; i++ {
		fmt.Printf("fib(%d) is: %d\n", i, fib(i))
	}
}

func fib(n int) (ans int) {
	if n <= 1 {
		ans = 1
	} else {
		ans = fib(n-1) + fib(n-2)
	}
	return
}
