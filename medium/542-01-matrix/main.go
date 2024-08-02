package main

import (
	"fmt"
	"math"
)

/*
Given an m x n binary matrix mat, return the distance of the nearest 0 for each cell.

The distance between two adjacent cells is 1.


Example 1:

	Input: mat = [[0,0,0],[0,1,0],[0,0,0]]
	Output: [[0,0,0],[0,1,0],[0,0,0]]

Example 2:

	Input: mat = [[0,0,0],[0,1,0],[1,1,1]]
	Output: [[0,0,0],[0,1,0],[1,2,1]]


Constraints:

	m == mat.length
	n == mat[i].length
	1 <= m, n <= 10^4
	1 <= m * n <= 10^4
	mat[i][j] is either 0 or 1.
	There is at least one 0 in mat.
*/

// Direction vectors for moving in the grid
var directions = [][]int{
	{-1, 0}, {1, 0}, {0, -1}, {0, 1},
}

// Function to check if a cell is within the matrix bounds
func inBounds(x, y, m, n int) bool {
	return x >= 0 && y >= 0 && x < m && y < n
}

// Function to find the distance of the nearest 0 for each cell in the binary matrix
func updateMatrix(mat [][]int) [][]int {
	m := len(mat)
	n := len(mat[0])

	// Initialize the result matrix with infinity for cells with 1
	// and enqueue all cells with 0
	queue := [][]int{}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if mat[i][j] == 0 {
				queue = append(queue, []int{i, j})
			} else {
				mat[i][j] = math.MaxInt32
			}
		}
	}

	// BFS to update distances
	for len(queue) > 0 {
		cell := queue[0]
		queue = queue[1:]
		x, y := cell[0], cell[1]

		for _, d := range directions {
			newX, newY := x+d[0], y+d[1]
			if inBounds(newX, newY, m, n) && mat[newX][newY] > mat[x][y]+1 {
				mat[newX][newY] = mat[x][y] + 1
				queue = append(queue, []int{newX, newY})
			}
		}
	}

	return mat
}

func main() {
	// Example 1
	mat1 := [][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}
	fmt.Println(updateMatrix(mat1)) // Output: [[0,0,0],[0,1,0],[0,0,0]]

	// Example 2
	mat2 := [][]int{{0, 0, 0}, {0, 1, 0}, {1, 1, 1}}
	fmt.Println(updateMatrix(mat2)) // Output: [[0,0,0],[0,1,0],[1,2,1]]
}
