package main

import (
	"fmt"
	"strings"
)

func main() {
	number := 20
	for i := 0; i <= number; i++ {
		if i == 20 {
			fmt.Println(repeatEmpty(number-i) + "#x" + repeatEmpty(i-1) + "|z" + repeatEmpty(i-1) + "#y")
			break
		}
		fmt.Println(repeatEmpty(number-i) + "#" + repeatEmpty(i) + "|" + repeatEmpty(i) + "#")

	}

	for i := 0; i <= number; i++ {
		fmt.Println(strings.Repeat(" ", number-i) + strings.Repeat("#", i*2))
	}

}

func repeatEmpty(i int) string {
	return strings.Repeat(" ", i)
}
