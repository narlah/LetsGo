package main

import "fmt"

func main() {
	x := 1
	f := func() bool {
		return x == 1
	}
	fmt.Println(f())

}
W