package main

import (
	f "fmt"
)

func main() {
	//var sudoku =  [][]byte{
	//	{'5', '3', '.',  '.', '7', '.',  '.', '.', '.'},
	//	{'6', '.', '.',  '1', '9', '5',  '.', '.', '.'},
	//	{'.', '9', '8',  '.', '.','.',  '.', '6', '.'},
	//
	//	{'8', '.', '.',  '.', '6', '.',  '.', '.', '3'},
	//	{'4', '.', '.',  '8', '.', '3',  '.', '.', '1'},
	//	{'7', '.', '.',  '.', '2', '.',  '.', '.', '6'},
	//
	//	{'.', '6', '.',  '.', '.', '.',  '2', '8', '.'},
	//	{'.', '.', '.',  '4', '1', '9',  '.', '.', '5'},
	//	{'.', '.', '.',  '.', '8', '.',  '.', '7', '9'}}

	var sudoku = [][]byte{
		{'.', '.', '4', '.', '.', '.', '6', '3', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
		{'5', '.', '.', '.', '.', '.', '.', '9', '.'},
		{'.', '.', '.', '5', '6', '.', '.', '.', '.'},
		{'4', '.', '3', '.', '.', '.', '.', '.', '1'},
		{'.', '.', '.', '7', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '5', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
	}

	f.Print(isValidSudoku(sudoku))
}

func isValidSudoku(board [][]byte) bool {
	sudokuByte := [9][9]byte{}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				sudokuByte[i][j] = 0
			} else {
				sudokuByte[i][j] = byte(board[i][j] - '0')
			}
		}
	}
	return isValidSudokuByte(sudokuByte);
}

func isValidSudokuByte(board [9][9]byte) bool {
	size := 9
	if len(board) != len(board[0]) || len(board) != size { //outside dimensions
		return false
	}
	for i := 0; i < size; i++ { //rows
		if !isThoseNineValid(board[i]) {
			return false
		}
		var current = [9]byte{}
		for j := 0; j < size; j++ { //cols
			current[j] = board[j][i]
		}
		if !isThoseNineValid(current) {
			return false
		}
	}
	var boxPositions = [][]byte{//boxes
		{0, 0}, {3, 0}, {6, 0},
		{0, 3}, {3, 3}, {6, 3},
		{0, 6}, {3, 6}, {6, 6}}
	for _, arr := range boxPositions {
		x, y := arr[0], arr[1]
		var current = [9]byte{} //just lazy to parametrise that
		current[0] = board[x][y]
		current[1] = board[x+1][y]
		current[2] = board[x+2][y]
		current[3] = board[x][y+1]
		current[4] = board[x][y+2]
		current[5] = board[x+1][y+1]
		current[6] = board[x+2][y+2]
		current[7] = board[x+2][y+1]
		current[8] = board[x+1][y+2]
		if !isThoseNineValid(current) {
			return false
		}
	}
	return true
}

func isThoseNineValid(ringWraiths [9]byte) bool { //Nine for Mortal Men doomed to die
	var checker = make([]bool, 10)
	for _, k := range ringWraiths {
		if k != 0 && checker[k] {
			return false
		}
		checker[k] = true
	}
	return true
}

func printSudoku(board [9][9]byte) {
	for _, z := range board {
		f.Printf("%d ", z)
		for _, b := range z {
			f.Print(b)
		}
		f.Print("\n")
	}
}
