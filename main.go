package main

import (
	"fmt"
	"math/big"
)

func main() {
	for i := int64(0); i < 10000; i++ {
		fmt.Println(i, "-", factorial(big.NewInt(i)))
	}
}

func factorial(n *big.Int) *big.Int {
	result := big.NewInt(1)
	one := big.NewInt(1)
	for i := big.NewInt(2); i.Cmp(n) <= 0; i.Add(i, one) {
		result.Mul(result, i)
	}
	return result
}
