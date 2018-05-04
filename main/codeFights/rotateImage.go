package main

import (
	"fmt"
	"math"
)

func main() {
	a := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	//a := [][]int{
	//	{1, 2, 3, 4},
	//	{5, 6, 7, 8},
	//	{9, 10, 11, 12},
	//	{13, 14, 15, 16},
	//}
	printNN(a)

	z := rotateImage(a)
	printNN(z)
}
func printNN(z [][]int) {
	for i := 0; i < len(z[0]); i++ {
		fmt.Println(z[i])
	}
	fmt.Println()
}

func rotateImage(a [][]int) [][]int {
	n := len(a[0])
	n2 := n / 2

	for i := 0; i < int(math.Ceil(float64(n2))); i++ {
		for j := i; j < n-1-i; j++ {

			rotate(a, i, j)
		}
	}
	return a
}
func rotate(a [][]int, i int, j int) {
	N := len(a[0])
	mai := N - 1 - i
	//a[i][j] //top
	//a[j][N-1-i] //right
	//a[N-1-i][N-1-j] //bottom
	//a[N-1-j][i] //left



	fmt.Printf(" %d %d mai: %d \n\n", i, j, mai)

	//store current cell in temp variable
	tmp := a[i][j]

	// left -> top
	a[i][j] = a[N-1-j][i]

	// bottom -> left
	a[N-1-j][i] = a[N-1-i][N-1-j]

	// right -> bottom
	a[N-1-i][N-1-j] = a[j][N-1-i]

	// assign temp to left
	a[j][N-1-i] = tmp
}
