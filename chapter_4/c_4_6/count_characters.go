package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s1 := "asSASA ddd dsjkdsjs dk"
	s2 := "asSASA ddd dsjkdsjsこん dk"

	fmt.Printf("bytes s1: %d, rune s1: %d\n", len(s1), utf8.RuneCountInString(s1))
	fmt.Printf("bytes s2: %d, rune s2: %d\n", len(s2), utf8.RuneCountInString(s2))
}
