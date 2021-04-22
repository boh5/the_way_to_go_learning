package main

import "fmt"

func main() {
	sl := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(Filter(sl, isOdd))
}

func Filter(sl []int, fn func(int) bool) (ret []int) {
	for _, val := range sl {
		if fn(val) {
			ret = append(ret, val)
		}
	}
	return
}

func isOdd(a int) bool {
	if a%2 == 0 {
		return true
	}
	return false
}
