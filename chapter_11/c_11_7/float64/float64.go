package float64

import (
	"fmt"
	"math/rand"
)

type Float64Array []float64

func (fa Float64Array) Len() int {
	return len(fa)
}

func (fa Float64Array) Less(i, j int) bool {
	return fa[i] < fa[j]
}

func (fa Float64Array) Swap(i, j int) {
	fa[i], fa[j] = fa[j], fa[i]
}

func (fa Float64Array) String() string {
	return List(fa)
}

func NewFloat64Array() *Float64Array {
	a := make(Float64Array, 25)
	return &a
}

func List(fa Float64Array) (ret string) {
	ret = "["
	for _, v := range fa {
		ret += " " + fmt.Sprint(v)
	}
	ret += " ]"
	return
}

func Fill(fa *Float64Array) {
	for i := 0; i < 10; i++ {
		(*fa)[rand.Intn(fa.Len())] = rand.Float64() * float64(rand.Intn(100))
	}
}
