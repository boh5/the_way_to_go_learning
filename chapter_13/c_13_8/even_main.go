package main

import (
	"fmt"

	"github.com/dilless/the_way_to_go_learning/chapter_13/c_13_8/even"
)

func main() {
	for i := 0; i < 101; i++ {
		fmt.Printf("Is the integer %d even? %v\n", i, even.Even(i))
	}
}
