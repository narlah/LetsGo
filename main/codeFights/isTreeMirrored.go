package main

import ("fmt"
	"strconv"
)

// Definition for binary tree:
type Tree struct {
 Value interface{}
 Left *Tree
 Right *Tree
}
func isTreeSymmetric(t *Tree) bool {
	if t == nil { return true}
	var leftSet = make([]*Tree, 0)
	var rightSet = make([]*Tree, 0)
	leftSet = append(leftSet, t.Left)
	rightSet = append(rightSet, t.Right)
	for len(leftSet)!=0 && len(rightSet)!=0{
		if !checkLR(leftSet, rightSet) {
			return false
		}
		leftSet = expandSet(leftSet)
		rightSet = expandSet(rightSet)
	}
	return true
}

func expandSet(set []*Tree) []*Tree{
	res := make([]*Tree, 0)
	for _,e := range set{
			if e == nil {continue}
			res = append(res, e.Left)
			res = append(res, e.Right)
		}
	return res
}
func checkLR(l []*Tree, r []*Tree) bool {
	ln := len(l)
	if len(l) != len(r) {
		return false
	}
	for i := 0; i < ln; i++ {
		if l[i]==nil && r[ln-i-1]==nil {continue}
		if l[i]==nil || r[ln-i-1]==nil || l[i].Value != r[ln-i-1].Value {
			return false
		}
	}
	return true
}

func main() {
	//root := &Tree{1,
	//&Tree{2,&Tree{3, nil , nil}, &Tree{4,nil , nil}},
	//	&Tree{2,&Tree{4, nil , nil}, &Tree{3,nil , nil}}};
	root := &Tree{1,
		&Tree{2,nil, &Tree{3,nil , nil}},
		&Tree{2,nil, &Tree{3,nil , nil}}};
	fmt.Printf("result : %s \n", strconv.FormatBool(isTreeSymmetric(root)))
}
/*
			 1
		   /   \
		  2      2
		 / \     / \
		3   4   4   3
	   / \  /\  /\  /\
	  5  6 7 8  8 7 6 5
//michael_c72
func isTreeSymmetric(t *Tree) bool {
    return helper(t, t)
}
*/
func helper (left *Tree, right *Tree) bool {
    if left == nil && right == nil {
        return true
    }

    if left == nil || right == nil || left.Value != right.Value {
        return false
    }

    return helper(left.Left, right.Right) && helper(left.Right, right.Left)
}

