package main

import (
	"fmt"

	float642 "github.com/dilless/the_way_to_go_learning/chapter_11/c_11_7/float64"
	"github.com/dilless/the_way_to_go_learning/chapter_11/c_11_7/sort"
)

func main() {
	fa := float642.NewFloat64Array()
	float642.Fill(fa)
	fmt.Println(*fa)
	sort.Sort(*fa)
	fmt.Println(*fa)
}
