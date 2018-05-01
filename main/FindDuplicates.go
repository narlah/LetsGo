package main

import (
	"math"
	"fmt"
)

func firstDuplicate(arr []int) int {
	size := len(arr)
	for i := 0; i < size; i++ {
		abs := int(math.Abs(float64(arr[i])))
		if arr[abs-1] < 0 {
			return abs
		}
		arr[abs-1] = -arr[abs-1]
	}
	return -1
}

func main() {
	arr := []int{2, 4, 3, 5, 1}
	fmt.Printf("result = %d",firstDuplicate(arr))
}
