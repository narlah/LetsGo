package main

import (
	"math"
	"fmt"
)

func main() {
	a1 := []int{1, 5, 6, 8, 9, 12, 15}
	a2 := []int{2, 7, 10, 11, 13, 14, 15}
	resultLen := len(a1) + len(a2)
	res := make([]int, resultLen)
	a1p, a2p := 0, 0
	for counter := 0; counter < resultLen; counter++ {
		addIT := 0
		vala1, vala2 := math.MaxInt8, math.MaxInt8
		if a1p < len(a1) {
			vala1 = a1[a1p]
		}
		if a2p < len(a2) {
			vala2 = a2[a2p]
		}

		if vala1 <= vala2 {
			addIT = a1[a1p]
			a1p++
		} else {
			addIT = a2[a2p]
			a2p++
		}
		res[counter] = addIT
	}
	fmt.Printf("merged %v", res)

}
