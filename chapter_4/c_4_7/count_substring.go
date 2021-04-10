package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string = "Hello, how is it going, Hugo?"
	var manyG = "gggggggggg"

	fmt.Printf("Number of H's in %s is: %d\n", str, strings.Count(str, "H"))

	fmt.Printf("Number of double g's in %s is: %d\n", manyG, strings.Count(manyG, "gg"))
}
