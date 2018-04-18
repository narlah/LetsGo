package main

import (
	"fmt"
	"strconv"
)

type DoubleLinkedRune struct {
	forward  *DoubleLinkedRune
	backward *DoubleLinkedRune
	rune     rune
}

func main() {
	root := &DoubleLinkedRune{nil, nil, '-'}
	runeArr := []rune{'a', 'b', 'c', 'd', 'e', 'f', 'g'}
	tempDlr := root
	last := root
	for _, r := range runeArr {
		newDLR := &DoubleLinkedRune{forward: nil, backward: nil, rune: r}
		tempDlr.backward = newDLR
		newDLR.forward = tempDlr
		tempDlr = newDLR
		fmt.Print(tempDlr)
		last = tempDlr
	}
	printList(root)

	tempDlr = root
	for n := 3; n > 0 && tempDlr != nil; n-- {
		tempDlr = tempDlr.backward
	}
	println(strconv.QuoteRune(last.rune))
	last.backward, root, tempDlr.backward = root, tempDlr.backward, nil
	printList(root)
}
func printList(root *DoubleLinkedRune) {
	fmt.Println("\n---------------------------")
	tempDlr := root
	for tempDlr != nil {
		fmt.Print(strconv.QuoteRune(tempDlr.rune), " ")
		tempDlr = tempDlr.backward
	}
}

func (d DoubleLinkedRune) String() string {
	return fmt.Sprintf(" [rune : %d  %d %d] ", d.rune, d.forward, d.backward)
}
