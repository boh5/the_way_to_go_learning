package main

import "fmt"

func main() {
	sl := []byte{'a', 'b', 'c'}
	fmt.Println(AppendByte(sl, []byte{'d', 'e', 'f'}...))
}

func AppendByte(slice []byte, data ...byte) []byte {
	m := len(slice)
	n := m + len(data)
	if n > cap(slice) {
		newSlice := make([]byte, (n+1)*2)
		copy(newSlice, slice)
		slice = newSlice
	}
	slice = slice[:n]
	copy(slice[m:n], data)
	return slice
}
