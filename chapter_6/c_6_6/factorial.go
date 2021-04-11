package main

import (
	"fmt"
	"math/big"
)

func main() {
	for i := 1; i < 33; i++ {
		fmt.Printf("fac(%d) is: %d\n", i, fac(big.NewInt(int64(i))))
	}
}

func fac(n *big.Int) (ans *big.Int) {
	if n.Int64() == 1 {
		return big.NewInt(1)
	} else {
		return n.Mul(n, fac(big.NewInt(n.Int64()-1)))
	}

}
