package main

import "fmt"

/*
Write a program to solve a Sudoku puzzle by filling the empty cells.

A sudoku solution must satisfy all of the following rules:

	1. Each of the digits 1-9 must occur exactly once in each row.
	2. Each of the digits 1-9 must occur exactly once in each column.
	3. Each of the digits 1-9 must occur exactly once in each of the 9 3x3 sub-boxes of the grid.
The '.' character indicates empty cells.


Example 1:

	Input: board = [["5","3",".",".","7",".",".",".","."],["6",".",".","1","9","5",".",".","."],[".","9","8",".",".",".",".","6","."],
					["8",".",".",".","6",".",".",".","3"],["4",".",".","8",".","3",".",".","1"],["7",".",".",".","2",".",".",".","6"],
					[".","6",".",".",".",".","2","8","."],[".",".",".","4","1","9",".",".","5"],[".",".",".",".","8",".",".","7","9"]]

	Output: [["5","3","4","6","7","8","9","1","2"],["6","7","2","1","9","5","3","4","8"],["1","9","8","3","4","2","5","6","7"],
			 ["8","5","9","7","6","1","4","2","3"],["4","2","6","8","5","3","7","9","1"],["7","1","3","9","2","4","8","5","6"],
			 ["9","6","1","5","3","7","2","8","4"],["2","8","7","4","1","9","6","3","5"],["3","4","5","2","8","6","1","7","9"]]

	Explanation: The input board is shown above and the only valid solution is shown below:


Constraints:

	board.length == 9
	board[i].length == 9
	board[i][j] is a digit or '.'.
	It is guaranteed that the input board has only one solution.
*/

func solveSudoku(board [][]byte) {
	solve(board)
}

func solve(board [][]byte) bool {
	// Find the next empty cell
	row, col, found := findEmptyCell(board)
	if !found {
		// No empty cell found, solution is complete
		return true
	}

	// Try placing each digit (1-9) in the empty cell
	for num := byte('1'); num <= '9'; num++ {
		if isValid(board, row, col, num) {
			board[row][col] = num
			if solve(board) {
				return true
			}
			board[row][col] = '.' // Backtrack
		}
	}

	return false
}

func findEmptyCell(board [][]byte) (int, int, bool) {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				return i, j, true
			}
		}
	}
	return -1, -1, false
}

func isValid(board [][]byte, row, col int, num byte) bool {
	// Check row
	for j := 0; j < 9; j++ {
		if board[row][j] == num {
			return false
		}
	}

	// Check column
	for i := 0; i < 9; i++ {
		if board[i][col] == num {
			return false
		}
	}

	// Check 3x3 sub-box
	startRow := (row / 3) * 3
	startCol := (col / 3) * 3
	for i := startRow; i < startRow+3; i++ {
		for j := startCol; j < startCol+3; j++ {
			if board[i][j] == num {
				return false
			}
		}
	}

	return true
}

func main() {
	board := [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'}, {'6', '.', '.', '1', '9', '5', '.', '.', '.'}, {'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'}, {'4', '.', '.', '8', '.', '3', '.', '.', '1'}, {'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'}, {'.', '.', '.', '4', '1', '9', '.', '.', '5'}, {'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}

	solveSudoku(board)

	for _, row := range board {
		fmt.Println(row)
	}
}
