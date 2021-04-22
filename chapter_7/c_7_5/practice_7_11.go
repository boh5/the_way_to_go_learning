package main

import "fmt"

func main() {
	sl := []int{1, 2, 3, 4, 5}
	fmt.Println(InsertIntSlice(sl, 100, 1))
}

func InsertIntSlice(sl []int, val, ix int) (ret []int) {
	ret = append(sl[:ix], append([]int{val}, sl[ix:]...)...)
	return
}
