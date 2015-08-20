package main

import (
	"fmt"
	"math/big"
)

func factorial(n int64) *big.Int {
	if n < 0 {
		return big.NewInt(1)
	}
	if n == 0 {
		return big.NewInt(1)
	}
	bigN := big.NewInt(n)
	return bigN.Mul(bigN, factorial(n-1))
}

func main() {
	fmt.Println(factorial(60))
}
