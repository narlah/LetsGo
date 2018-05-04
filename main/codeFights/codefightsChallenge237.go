package main

import (
	"fmt"
)

func main() {
	fmt.Printf("what the ... really : %d", str237("The Cameroon country code 237 will allow you to call Cameroon from another country. Cameroon telephone code 237 is dialed after the IDD. Cameroon international dialing 237 is followed by an area code."));
}

//da fak, i don't even know where to start ...
func str237(S string) int {
	D := [238]int{1}
	for _, c := range S {
		var x int32 = 238
		for x > c {
			x--
			D[x] += D[x-c]
		}
	}
	fmt.Println(D)
	return D[237]
}
