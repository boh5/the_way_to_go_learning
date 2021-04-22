package main

import "fmt"

func main() {
	sl := []int{1, 2, 3, 4, 5}
	fmt.Println(InsertIntSlice(sl, 100, 1))
}

func InsertIntSlice(sl []int, val, ix int) (ret []int) {
	var before, after []int
	before = append(before, sl[:ix]...)
	after = append(after, sl[ix:]...)
	ret = append(append(before, val), after...)
	return
}
