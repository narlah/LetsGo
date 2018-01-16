package main

import (
	f "fmt"
	m "math"
)

func main() {

	f.Printf("fib number for 10 : %d \n ", fibonaciInt(10))
	f.Printf("fib number for 10 : %d \n ", int(binetFormulla(10)))

	f.Printf("fib number for 15 : %d \n ", fibonaciInt(15))
	f.Printf("fib number for 10 : %d \n ", int(binetFormulla(15)))

	f.Printf("fib number for 100 : %d \n ", fibonaciInt(50))
	f.Printf("fib number for 10 : %d \n ", int(binetFormulla(50)))
}

func fibonaciInt(n int) int {
	if n < 2 {
		return n
	} else {
		var sum int = 0
		var nMinus1 int = 1
		var nMinus2 int = 1
		for i := 2; i < n; i++ {
			sum = nMinus1 + nMinus2;
			nMinus2 = nMinus1;
			nMinus1 = sum;
		}
		return sum;
	}
}

func binetFormulla(n int) float64 {
	sqrt5 := m.Sqrt(5)
	finNum := (1 / sqrt5) * ( m.Pow((1+sqrt5)/2, float64(n)) - m.Pow((1-sqrt5)/2, float64(n)) )
	return finNum
}
