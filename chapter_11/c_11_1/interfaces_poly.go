package main

import "fmt"

type Shaper1 interface {
	Area() float32
}

type Square1 struct {
	side float32
}

func (sq *Square1) Area() float32 {
	return sq.side * sq.side
}

type Rectangle1 struct {
	length, width float32
}

func (r Rectangle1) Area() float32 {
	return r.length * r.width
}

func main() {
	r := Rectangle1{5, 3} // Area() of Rectangle needs a value
	q := &Square1{5}      // Area() of Square needs a pointer
	// shapes := []Shaper{Shaper(r), Shaper(q)}
	// or shorter
	shapes := []Shaper1{r, q}
	fmt.Println("Looping through shapes for area ...")
	for n, _ := range shapes {
		fmt.Println("Shape details: ", shapes[n])
		fmt.Println("Area of this shape is: ", shapes[n].Area())
	}
}
