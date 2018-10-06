package main

import "fmt"

func main() {
	str := ""
	for i := 1; i <= 100; i++ {
		if i%5 == 0 {
			str += "Fizz"
		}
		if i%3 == 0 {
			str += "Buzz"
		}
		fmt.Printf("%d : %s\n", i, str)
		str = ""
	}

}
