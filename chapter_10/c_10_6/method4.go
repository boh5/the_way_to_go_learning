package main

import (
	"fmt"
	"math"
)

type Point1 struct {
	x, y float64
}

type NamedPoint1 struct {
	Point1
	name string
}

func main() {
	n := &NamedPoint1{Point1{3, 4}, "Pythagoras"}
	fmt.Println(n.Abs())
}

func (p *Point1) Abs() float64 {
	return math.Sqrt(p.x*p.x + p.y*p.y)
}

func (n NamedPoint1) Abs() float64 {
	return n.Point1.Abs() * 100
}
