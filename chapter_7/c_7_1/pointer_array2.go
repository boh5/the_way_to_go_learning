package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 3; i++ {
		fp2(&[3]int{i, i * i, i * i * i})
	}
}

func fp2(a *[3]int) {
	fmt.Println(a)
}
