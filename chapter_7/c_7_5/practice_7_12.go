package main

import "fmt"

func main() {
	sl := []string{"a", "b", "c", "d", "e", "f", "g"}
	fmt.Println(RemoveStringSlice(sl, 1, 3))
}

func RemoveStringSlice(sl []string, start, end int) (ret []string) {
	ret = append(sl[:start], sl[end+1:]...)
	return
}
