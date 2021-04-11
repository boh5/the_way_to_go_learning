package main

import "fmt"

func main() {
	f(1, 2, 'a', "Bc")
}

func f(s ...interface{}) {
	for _, v := range s {
		fmt.Println(v)
	}
}
