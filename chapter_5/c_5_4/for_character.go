package main

import (
	"fmt"
)

func main() {
	forInFor()
	forStringConcat()
}

func forInFor() {
	for i := 0; i < 25; i++ {
		for j := 0; j <= i; j++ {
			fmt.Printf("G")
		}
		fmt.Println()
	}
}

func forStringConcat() {
	str := "G"
	for i := 0; i < 25; i++ {
		fmt.Println(str)
		str += "G"
	}
}
