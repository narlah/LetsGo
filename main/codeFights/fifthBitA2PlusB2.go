package main

import (
	"fmt"
	"strconv"
)

func main() {
	check(826921131, 48669151, 1) //1
	check(380307315, 524023846, 0)
	check(709494765, 480216299, 1)

	check(717757366, 213356038, 0)
	check(855757028, 879543000, 0)
	check(251129265, 673395586, 1)
	check(105198459, 805822999, 1)

}

//So apparently fifthBit means the bit under 2^5, so 2^0-2^4(total of five bits) has been removed and then we get the fifth bit.
func fifthBitA2PlusB2(a int, b int) int {
	var bits uint = 5
	n := int64((a * a ) + (b * b))
	fmt.Printf("%d %s %s                        ", n, strconv.FormatInt(n, 2), strconv.FormatInt(n>>bits, 2))
	return (((a * a ) + (b * b)) >> bits ) & 1
}

func check(a int, b int, expected int) {
	res := fifthBitA2PlusB2(a, b)
	flag := "False"
	if expected == res {
		flag = "True"
	}
	fmt.Printf("a: %d | b: %d | expected: %d == %d %s\n ", a, b, expected, res, flag)
}
