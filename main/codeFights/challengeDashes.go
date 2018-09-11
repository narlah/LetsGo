package main

import (
	"fmt"
)

func main() {
	arr := dashes(12)
	for i := 0; i < len(arr); i++ {
		fmt.Printf("[%s]\n",arr[i])
	}
}

func dashes(n int) []string {
	total := 2*n - 1
	var arr = make([]string, total)
	for i := 0; i <= total/2; i++ {
		s := ""
		m := ""
		for l := 0; l < n-i-1; l++ {
			s = s + " "
		}
		for p := 1; p <= 2*i+1 ; p++ {
			if p%2 == 0 {
				m += "|"
			} else {
				m += "-"
			}
		}
		arr[i] = s + m + s
		arr[total-i-1] = s + m + s
	}

	return arr
	// 1-1 2-3 3-5 4-7
}
