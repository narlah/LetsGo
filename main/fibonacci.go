package main

import (
	f "fmt"
	m "math"
)

func main() {

    f.Printf("int : fib number for 10 : %d \n ", fibonaciInt(10))
    f.Printf("formula : fib number for 10 : %d \n ", int(binetFormulla(10)))

    f.Printf("int : fib number for 15 : %d \n ", fibonaciInt(15))
	f.Printf("formula fib number for 15 : %d \n ", int(binetFormulla(15)))

	f.Printf("int fib number for 50 : %d \n ", fibonaciInt(50))
	f.Printf("formula fib number for 50 : %d \n ", int64(binetFormulla(50)))
}

func fibonaciInt(n int) int {
	if n < 2 {
		return n
	}
	var sum int
	nMinus1 := 1
	nMinus2 := 1
	for i := 2; i < n; i++ {
		sum = nMinus1 + nMinus2
		nMinus2 = nMinus1
		nMinus1 = sum
	}
	return sum
}

func binetFormulla(num int) int {
	sqrt5 := float64(m.Sqrt(5))
    n := float64(num)
    return int((m.Pow((1+sqrt5)/2, n) - m.Pow((1-sqrt5)/2, (n))) / sqrt5 +0.5)

}
