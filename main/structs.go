package main

import (
	"fmt"
)

type stt struct {
	a int
	b int
	s string
}

func main() {
	var sttSlice []stt
	for i := 0; i < 10; i++ {
		struct1 := stt{i, i*10, "a"}
		sttSlice = append(sttSlice, struct1)
	}

	fmt.Println(sttSlice)
}
