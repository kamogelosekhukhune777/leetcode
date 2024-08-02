package main

import (
	"fmt"
)

/*
The n-queens puzzle is the problem of placing n queens on an n x n chessboard
such that no two queens attack each other.

Given an integer n, return all distinct solutions to the n-queens puzzle. You
may return the answer in any order.

Each solution contains a distinct board configuration of the n-queens' placement, where
'Q' and '.' both indicate a queen and an empty space, respectively.


Example 1:

	Input: n = 4
	Output: [[".Q..","...Q","Q...","..Q."],["..Q.","Q...","...Q",".Q.."]]
	Explanation: There exist two distinct solutions to the 4-queens puzzle as shown above

Example 2:

	Input: n = 1
	Output: [["Q"]]


Constraints:

1 <= n <= 9
Seen
*/

func solveNQueens(n int) [][]string {
	var result [][]string
	board := make([][]rune, n)
	for i := range board {
		board[i] = make([]rune, n)
		for j := range board[i] {
			board[i][j] = '.'
		}
	}
	backtrack(board, 0, &result)
	return result
}

func backtrack(board [][]rune, row int, result *[][]string) {
	if row == len(board) {
		// Convert board to the desired output format
		var solution []string
		for _, row := range board {
			solution = append(solution, string(row))
		}
		*result = append(*result, solution)
		return
	}

	for col := 0; col < len(board); col++ {
		if isValid(board, row, col) {
			board[row][col] = 'Q'
			backtrack(board, row+1, result)
			board[row][col] = '.'
		}
	}
}

func isValid(board [][]rune, row, col int) bool {
	// Check the column
	for i := 0; i < row; i++ {
		if board[i][col] == 'Q' {
			return false
		}
	}

	// Check the major diagonal
	for i, j := row-1, col-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if board[i][j] == 'Q' {
			return false
		}
	}

	// Check the minor diagonal
	for i, j := row-1, col+1; i >= 0 && j < len(board); i, j = i-1, j+1 {
		if board[i][j] == 'Q' {
			return false
		}
	}

	return true
}

func main() {
	n := 4
	solutions := solveNQueens(n)
	for _, solution := range solutions {
		for _, line := range solution {
			fmt.Println(line)
		}
		fmt.Println()
	}
}
