package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {
	ans, err := MySqrt(-11)
	if err != nil {
		fmt.Println("Error!", ans, err)
	} else {
		fmt.Println("OK.", ans, err)
	}

	ans, err = MySqrt2(25)
	if err != nil {
		fmt.Println("Error!", ans, err)
	} else {
		fmt.Println("OK.", ans, err)
	}
}

func MySqrt(x float64) (float64, error) {
	if x < 0 {
		return math.NaN(), errors.New("input less than 0")
	}
	return math.Sqrt(x), nil
}

func MySqrt2(x float64) (ans float64, err error) {
	if x < 0 {
		ans, err = math.NaN(), errors.New("input less than 0")
	} else {
		ans, err = math.Sqrt(x), nil
	}
	return
}
