package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan float64)
	n := 50000
	go put(ch, n)
	var ans float64
	for f := range ch {
		ans += f
	}
	fmt.Println(ans)
	end := time.Now()
	fmt.Println("time:", end.Sub(start))
}

func put(ch chan float64, n int) {
	for i := 0; i < n+1; i++ {
		ch <- 4 * math.Pow(-1, float64(i)) / float64(2*i+1)
	}
	close(ch)
}
