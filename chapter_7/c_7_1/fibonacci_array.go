package main

import "fmt"

func main() {
	for i := 1; i < 51; i++ {
		fmt.Println(i, fib(i))
	}
}

func fib(n int) (a int) {
	var fibArray [50]int
	for i := 0; i < n; i++ {
		if i == 0 || i == 1 {
			fibArray[i] = 1
		} else {
			fibArray[i] = fibArray[i-1] + fibArray[i-2]
		}
	}

	a = fibArray[n-1]
	return
}
