package main

import "fmt"

func main() {
	fmt.Println(f1(3, 4))
	fmt.Println(f2(3, 2))
}

func f1(x, y int) (int, int, int) {
	sum := x + y
	pro := x * y
	diff := x - y
	return sum, pro, diff
}

func f2(x, y int) (sum, pro, diff int) {
	sum = x + y
	pro = x * y
	diff = x - y
	return
}
