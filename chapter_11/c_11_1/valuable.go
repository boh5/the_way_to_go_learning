package main

import "fmt"

type stopPosition struct {
	ticker     string
	sharePrice float32
	count      float32
}

type car struct {
	make  string
	model string
	price float32
}

type valuable interface {
	getValue() float32
}

func main() {
	var o valuable = stopPosition{"GOOG", 577.20, 4}
	showValue(o)
	o = car{"BMW", "M3", 66500}
	showValue(o)
}

func (s stopPosition) getValue() float32 {
	return s.sharePrice * s.count
}

func (c car) getValue() float32 {
	return c.price
}

func showValue(asset valuable) {
	fmt.Printf("Value of the asset is %f\n", asset.getValue())
}
