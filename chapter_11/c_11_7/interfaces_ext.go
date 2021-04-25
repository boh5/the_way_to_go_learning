package main

import "fmt"

type AreaInterface interface {
	Area() float64
}

type PeriInterface interface {
	Perimeter() int
}

type Triangle struct {
	a int
	h int
}

type Square struct {
	a int
}

func main() {
	t := &Triangle{1, 2}
	fmt.Println(t.Area())
	s := &Square{4}
	fmt.Println(s.Perimeter())
}

func (t *Triangle) Area() float64 {
	return 0.5 * float64(t.a) * float64(t.h)
}

func (s Square) Perimeter() int {
	return 4 * s.a
}
