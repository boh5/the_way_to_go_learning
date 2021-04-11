package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	fmt.Printf(nonAscii("abcd我的", '?'))
}

func nonAscii(orig string, part rune) string {
	newS := strings.Map(func(r rune) rune {
		if !unicode.Is(unicode.ASCII_Hex_Digit, r) {
			return part
		}
		return r
	}, orig)
	return newS
}
