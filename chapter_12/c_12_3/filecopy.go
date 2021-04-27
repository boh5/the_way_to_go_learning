package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	count, err := CopyFile("D:\\Files\\codes\\golang\\the_way_to_go_learning\\chapter_12\\c_12_3\\target.txt",
		"D:\\Files\\codes\\golang\\the_way_to_go_learning\\chapter_12\\c_12_3\\source.txt")
	fmt.Println(count)
	fmt.Println(err)
	fmt.Println("Copy done!")
}

func CopyFile(dstName, srcName string) (written int64, err error) {
	src, err := os.Open(srcName)
	if err != nil {
		return
	}
	defer src.Close()

	dst, err := os.Create(dstName)
	if err != nil {
		return
	}
	defer dst.Close()

	return io.Copy(dst, src)
}
