package main

import "fmt"

func main() {
	fmt.Println(revertString("Google"))
}

func revertString(s string) (ret string) {
	sb := []byte(s)
	for i := 0; i < len(sb)/2; i++ {
		sb[i], sb[len(sb)-1-i] = sb[len(sb)-1-i], sb[i]
	}
	ret = string(sb)
	return
}
