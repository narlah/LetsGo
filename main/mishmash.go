package main

import (
	"strconv"
	"fmt"
)

func main() {
	fmt.Print(mishMash(4))
}

func mishMash(num int) string {
	div2 := (num % 2) == 0
	div6 := (num % 6) == 0
	if div2 && div6 {
		return "MishMash"
	} else if div2 {
		return "Mish"
	} else if div6 {
		return "Mash"
	} else {
		return strconv.Itoa(num)
	}
}
