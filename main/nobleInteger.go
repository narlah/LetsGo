package main

import (
    "fmt"
    "sort"
)


func solve(A []int )  (int) {
    sort.Ints(A)

    for i:=0;i<len(A);i++ {
        if A[i] == 0 && i == len(A)-1 {
            return 1
        }
        if A[i] == len(A)-i-1 && i != len(A)-1 && A[i] != A[i+1]{
            return 1
        }
    }
    return -1
}

func main(){
    arr := []int{6,1}
    fmt.Printf("result %d " , solve(arr))
}
