package main

import (
	"flag"
	"fmt"
	"math"
)

type polar struct {
	r float64
	a float64
}

type cartesian struct {
	x, y float64
}

func main() {
	r := flag.Float64("r", 0.0, "半径")
	a := flag.Float64("a", 0.0, "角度")
	flag.Parse()
	p := polar{*r, *a}
	ch1 := make(chan polar)
	ch2 := make(chan cartesian)
	go func() {
		ch1 <- p
	}()
	go func() {
		pIn := <-ch1
		x := pIn.r * math.Cos(pIn.a)
		y := pIn.r * math.Sin(pIn.a)
		ch2 <- cartesian{x, y}

	}()
	fmt.Println(<-ch2)
}
