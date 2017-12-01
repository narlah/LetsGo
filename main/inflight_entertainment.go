package main

import (
	f "fmt"
)

func main() {
	movies := [...]int{2, 2, 6, 8, 9, 11, 12, 25, 26, 46}
	var flightLen int = 11
	f.Printf("result is : %s", checkInFlight(movies, flightLen))
}

func checkInFlight(movies [10]int, flightLen int) bool {
	hashSet := map[int]bool{}
	for _, movieLenght := range movies {
		reqLen := flightLen - movieLenght
		_, ok := hashSet[reqLen]
		//f.Printf("log %d ,  %d \n", reqLen, ok)
		if ok {
			return true
		}
		hashSet[movieLenght] = true
	}
	return false
}
