package main

import "fmt"

func main() {
	sl := []byte{'a', 'b', 'c'}
	fmt.Println(Append(sl, []byte{'d', 'e'}))
	fmt.Println(Append(sl, []byte{'d', 'e', 'f', 'g'}))
}

func Append(slice, data []byte) []byte {
	lenS := len(slice)
	lenD := len(data)
	totalLen := lenS + lenD
	if cap(slice) < totalLen {
		tmp := make([]byte, totalLen)
		for i := 0; i < lenS; i++ {
			tmp[i] = slice[i]
		}
		slice = tmp
	}
	for i := 0; i < lenD; i++ {
		slice[lenS+i] = data[i]
	}
	return slice
}
