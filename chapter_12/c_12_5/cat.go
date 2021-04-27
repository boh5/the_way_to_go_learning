package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	flag.Parse()
	if flag.NArg() == 0 {
		cat(bufio.NewReader(os.Stdin))
	}
	for i := 0; i < flag.NArg(); i++ {
		f, err := os.Open(flag.Arg(i))
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s: error reading from %s:%s\n", os.Args[0], flag.Arg(i), err.Error())
			continue
		}
		cat(bufio.NewReader(f))
		f.Close()
	}
}

func cat(r *bufio.Reader) {
	i := 0
	for {
		i += 1
		buf, err := r.ReadBytes('\n')
		fmt.Fprintf(os.Stdout, "%d: %s", i, buf)
		if err == io.EOF {
			break
		}
	}
	return
}
