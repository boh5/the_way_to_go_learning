package main

import (
	"fmt"

	"github.com/dilless/the_way_to_go_learning/chapter_9/c_9_5/pack1"
)

func main() {
	var test1 string
	test1 = pack1.ReturnStr()
	fmt.Printf("ReturnStr from package1: %s\n", test1)
	fmt.Printf("Integer from package1: %d\n", pack1.Pack1Int)
}
