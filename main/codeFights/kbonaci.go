package main

import (
	"fmt"
	m "math/big"
	. "container/list"
)

//1 ≤ n ≤ 10^5
//1 ≤ k ≤ 10^3
func main() {
	fmt.Println(kbonacci(3, 4)) //5
}

func kbonacci(k int, n int) string {
	if n <= k {
		return string(1)
	}
	listR := new(List)

	for i := 0; i < k; i++ {
		listR.PushBack(m.NewInt(1))
	}
	//// Iterate through list and print its contents.
	//for e := listR.Front(); e != nil; e = e.Next() {
	//	fmt.Println(e.Value)
	//}

	var res = m.NewInt(0)
	for z := 0; z < n-k; z++ {
		res = m.NewInt(0).Sub(parse(listR.Front().Value), res).Add(parse(listR.Back().Value), res)
		listR.Remove(listR.Front())
		listR.PushBack(res)
	}
	return res.Text(10)
}

func parse(in m.Int) *m.Int{
	return &in
}
