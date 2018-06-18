package main

import (
"fmt"
"math/big"
)

func main() {
	verybig := big.NewInt(1)
	ten := big.NewInt(5)
	for i:=0; i<1000; i++ {
		verybig.Mul(verybig, ten)
	}
	fmt.Println(verybig)
}