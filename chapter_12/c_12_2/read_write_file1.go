package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	inputFile := "chapter_12/c_12_2/products.txt"
	outputFile := "chapter_12/c_12_2/products_copy.txt"
	buf, err := ioutil.ReadFile(inputFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "File Error: %s\n", err)
		panic(err.Error())
	}
	fmt.Printf("%s\n", string(buf))
	err = ioutil.WriteFile(outputFile, buf, 0644)
	if err != nil {
		panic(err.Error())
	}
}
