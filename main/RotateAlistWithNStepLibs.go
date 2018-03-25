package main

import (
	"fmt"
	. "container/list"
	"strconv"
)

func main() {
	listR := new(List)

	runeArr := []rune{'a', 'b', 'c', 'd', 'e', 'f', 'g'}

	for _, r := range runeArr {
		listR.PushBack(r)
	}
	cycle(listR, 3)
	fmt.Println(listR.Len())
	for e := listR.Front(); e != nil; e = e.Next() {
		fmt.Print(strconv.QuoteRune(e.Value.(rune)))
	}

}

func cycle(l *List, n int) {
	for k := 0; k < n%l.Len(); k++ {
		l.MoveToFront(l.Back())
	}
}
