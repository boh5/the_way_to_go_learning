package main

import "fmt"

func main() {
	for i := 1; i <= 20; i++ {
		fmt.Println(fib(i))
	}
}

func fib(n int) (ret []int) {
	ret = make([]int, n)
	for i := 0; i < n; i++ {
		if i == 0 || i == 1 {
			ret[i] = 1
		} else {
			ret[i] = ret[i-1] + ret[i-2]
		}
	}
	return
}
