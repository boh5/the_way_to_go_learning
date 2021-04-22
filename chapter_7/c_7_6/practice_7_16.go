package main

import "fmt"

func main() {
	fmt.Println(BubbleSort([]int{1, 2, 3, 4, 5, 7, 6}))
}

func BubbleSort(sl []int) []int {
	for i := 0; i < len(sl); i++ {
		swapped := false
		for j := 0; j < len(sl)-1; j++ {
			if sl[j] > sl[j+1] {
				sl[j], sl[j+1] = sl[j+1], sl[j]
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
	return sl
}
