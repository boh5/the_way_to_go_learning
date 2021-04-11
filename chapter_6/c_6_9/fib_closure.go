package main

import "fmt"

func main() {
	f := fib()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}

func fib() func() int {
	x0 := 1
	x1 := 1
	return func() int {
		t := x0
		x0, x1 = x1, x1+x0
		return t
	}
}
