package main

import "fmt"

//
// Definition for binary tree:
//type Tree struct {
//  Value interface{}
//  Left *Tree
//  Right *Tree
//}
func hasPathWithGivenSum(t *Tree, s int) bool {
	if t == nil {return false}
	val := s-t.Value.(int)
	if val == 0 && t.Left == nil && t.Right == nil { return true }
	return hasPathWithGivenSum (t.Left, val ) || hasPathWithGivenSum (t.Right, val)
}
func main() {
	t := &Tree{5, nil, nil}
	fmt.Println(hasPathWithGivenSum(t,5))
}
