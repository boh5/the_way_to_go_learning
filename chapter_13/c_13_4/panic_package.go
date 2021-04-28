package main

import (
	"fmt"

	"github.com/dilless/the_way_to_go_learning/chapter_13/c_13_4/parse"
)

func main() {
	examples := []string{
		"1 2 3 4 5",
		"100 50 25 12.5 6.25",
		"2 + 2 = 4",
		"1st class",
		"",
	}

	for _, ex := range examples {
		fmt.Printf("Parsing %q:\n ", ex)
		nums, err := parse.Parse(ex)
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println(nums)
	}
}
