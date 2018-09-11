package main

import (
	"fmt"
	"sort"
)

func diceRolls(dice [][]int, sum int) int {
	diceLen := len(dice)
	arrNs := make([]int, diceLen)
	counter := 0
	for index, arr := range dice {
		x := sort.SearchInts(arr, sum)
		for _, y := range arr[x:] {
			if arr[x] != y {
				break
			}
			x++
		}
		arrNs[index] = x
	}
	for i, el := range dice[diceLen-1] {
		dice[diceLen-1][i] = sum - el
	}
	var fu func(currDiceIndex int, currSum int)
	fu = func(currDiceIndex int, currSum int) {
		if currDiceIndex == -1 && currSum == 0 {
			counter++
			return
		}
		if currDiceIndex < 0 {

			return
		}
		for i := arrNs[currDiceIndex] - 1; i >= 0; i-- {
			newSum := currSum - dice[currDiceIndex][i]
			if newSum < 0 {
				continue
			}
			fu(currDiceIndex-1, newSum)
		}
	}
	for _, el := range dice[diceLen-1] {
		fu(diceLen-2, el)
	}
	return counter
}

/*
11/16
func diceRolls(dice [][]int, sum int) int {
	diceLen := len(dice)
	arrNs := make([]int, diceLen)
	counter := 0
	for index, arr := range dice {
		x := sort.SearchInts(arr, sum)
		for _,y := range arr[x:] {
			if arr[x]!= y  { break }
			x++
		}
		arrNs[index] = x
	}
	var fu func(currDiceIndex int, currSum int)
	fu = func(currDiceIndex int, currSum int) {
		if currDiceIndex == -1 && currSum == sum {
			counter++
			return
		}
		if currDiceIndex < 0 {
			return
		}
		for i := arrNs[currDiceIndex] - 1; i >= 0; i-- {
			newSum := currSum + dice[currDiceIndex][i]
			if newSum > sum {
				continue
			}
			fu(currDiceIndex-1, newSum)
		}
	}
	fu(diceLen-1, 0)
	return counter
}
*/
/*
	MAJOR FAIL:P
	Always read the whole requirements, i though its just 2 dices ...
	apparently they can be a 1000 ... man , 4s execute time would be hard

	x := len(dice[0])
	y := len(dice[1])
	var smaller *[]int
	var bigger *[]int
	if x >= y {
		smaller = &dice[1]
		bigger = &dice[0]
	} else {
		smaller = &dice[0]
		bigger = &dice[1]

	}
	m := make(map[int]bool)
	for _, element := range *bigger {
		m[element] = true
	}
	counter := 0
	for _, smallerElement := range *smaller {
		if smallerElement < sum  && m[sum - smallerElement]{
			counter++
		}
	}
	return counter
*/

func main() {
	//dices := [][]int{{12,19,20,25,33,43,62,68,72,78,88,92,98},
	//	{0,1,4,11,22,37,45,53,74,79,82,83,89},
	//	{12,17,37,61,80,98},
	//	{14,15,16,24,46,67,77,79,80,82,83},
	//	{3,13,28,50,62,80,81,83,98}}
	//fmt.Println(diceRolls(dices, 45))
	//dice := [][]int{{1, 2, 3, 4, 5, 6}, {1, 2, 3, 4, 5, 6},}
	//fmt.Println(diceRolls(dice, 9))
	//dice := [][]int{{0, 0, 0, 0, 0, 0},}
	//fmt.Println(diceRolls(dice, 0))
	var dice [101][101]int
	for i := 0; i <= 100; i++ {
		for j := 0; j <= 100; j++ {
			dice[i][j] = j * i
		}
	}
	fmt.Println(diceRolls(dice[:][:], 12))
}
