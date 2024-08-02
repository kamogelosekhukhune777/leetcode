package main

import "fmt"

/*
An image is represented by an m x n integer grid image where image[i][j] represents
the pixel value of the image.

You are also given three integers sr, sc, and color. You should perform a flood fill on the
image starting from the pixel image[sr][sc].

To perform a flood fill, consider the starting pixel, plus any pixels connected 4-directionally
to the starting pixel of the same color as the starting pixel, plus any pixels connected 4-directionally
to those pixels (also with the same color), and so on. Replace the color of all of the
aforementioned pixels with color.

Return the modified image after performing the flood fill.



Example 1:

	Input: image = [[1,1,1],[1,1,0],[1,0,1]], sr = 1, sc = 1, color = 2
	Output: [[2,2,2],[2,2,0],[2,0,1]]
	Explanation: From the center of the image with position (sr, sc) = (1, 1) (i.e., the red pixel), all pixels connected by a path of the same color as the starting pixel (i.e., the blue pixels) are colored with the new color.
	Note the bottom corner is not colored 2, because it is not 4-directionally connected to the starting pixel.

Example 2:

	Input: image = [[0,0,0],[0,0,0]], sr = 0, sc = 0, color = 0
	Output: [[0,0,0],[0,0,0]]
	Explanation: The starting pixel is already colored 0, so no changes are made to the image.


Constraints:

	m == image.length
	n == image[i].length
	1 <= m, n <= 50
	0 <= image[i][j], color < 2^16
	0 <= sr < m
	0 <= sc < n

*/

func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	originalColor := image[sr][sc]
	if originalColor != newColor {
		dfs(image, sr, sc, originalColor, newColor)
	}
	return image
}

func dfs(image [][]int, r, c, color, newColor int) {
	// Check for out of bounds or if the current pixel is not the original color
	if r < 0 || r >= len(image) || c < 0 || c >= len(image[0]) || image[r][c] != color {
		return
	}

	// Change the color of the current pixel
	image[r][c] = newColor

	// Recursively apply the same process to the 4-directionally connected pixels
	dfs(image, r+1, c, color, newColor) // Down
	dfs(image, r-1, c, color, newColor) // Up
	dfs(image, r, c+1, color, newColor) // Right
	dfs(image, r, c-1, color, newColor) // Left
}

func main() {
	// Example 1
	image1 := [][]int{{1, 1, 1}, {1, 1, 0}, {1, 0, 1}}
	sr1, sc1, color1 := 1, 1, 2
	fmt.Println(floodFill(image1, sr1, sc1, color1)) // Output: [[2,2,2],[2,2,0],[2,0,1]]

	// Example 2
	image2 := [][]int{{0, 0, 0}, {0, 0, 0}}
	sr2, sc2, color2 := 0, 0, 0
	fmt.Println(floodFill(image2, sr2, sc2, color2)) // Output: [[0,0,0],[0,0,0]]
}
