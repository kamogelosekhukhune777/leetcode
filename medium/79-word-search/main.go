package main

import "fmt"

/*
Given an m x n grid of characters board and a string word, return true if word exists in the grid.

The word can be constructed from letters of sequentially adjacent cells, where adjacent cells are horizontally
or vertically neighboring. The same letter cell may not be used more than once.


Example 1:

	Input: board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "ABCCED"
	Output: true

Example 2:

	Input: board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "SEE"
	Output: true

Example 3:

	Input: board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "ABCB"
	Output: false


Constraints:

	m == board.length
	n = board[i].length
	1 <= m, n <= 6
	1 <= word.length <= 15
	board and word consists of only lowercase and uppercase English letters.


Follow up: Could you use search pruning to make your solution faster with a larger board?
*/

func exist(board [][]byte, word string) bool {
	for i := range board {
		for j := range board[i] {
			if backtrack(board, word, i, j, 0) {
				return true
			}
		}
	}
	return false
}

func backtrack(board [][]byte, word string, i, j, index int) bool {
	// Check if the current index equals the length of the word
	if index == len(word) {
		return true
	}
	// Check boundaries and character match
	if i < 0 || i >= len(board) || j < 0 || j >= len(board[i]) || board[i][j] != word[index] {
		return false
	}

	// Temporarily mark the cell as visited
	temp := board[i][j]
	board[i][j] = '#'

	// Explore all four directions
	if backtrack(board, word, i+1, j, index+1) ||
		backtrack(board, word, i-1, j, index+1) ||
		backtrack(board, word, i, j+1, index+1) ||
		backtrack(board, word, i, j-1, index+1) {
		return true
	}

	// Backtrack: reset the cell's value
	board[i][j] = temp

	return false
}

func main() {
	board := [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}
	word := "ABCCED"
	fmt.Println(exist(board, word)) // Output: true

	word = "SEE"
	fmt.Println(exist(board, word)) // Output: true

	word = "ABCB"
	fmt.Println(exist(board, word)) // Output: false
}
