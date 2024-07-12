package main

import (
	"container/heap"
	"fmt"
)

/*
Given an n x n matrix where each of the rows and columns is sorted in ascending order, return
the kth smallest element in the matrix.

Note that it is the kth smallest element in the sorted order, not the kth distinct element.

You must find a solution with a memory complexity better than O(n^2).


Example 1:

	Input: matrix = [[1,5,9],[10,11,13],[12,13,15]], k = 8
	Output: 13
	Explanation: The elements in the matrix are [1,5,9,10,11,12,13,13,15], and the 8th smallest
	number is 13
Example 2:

	Input: matrix = [[-5]], k = 1
	Output: -5


Constraints:

	n == matrix.length == matrix[i].length
	1 <= n <= 300
	-109 <= matrix[i][j] <= 109
	All the rows and columns of matrix are guaranteed to be sorted in non-decreasing order.
	1 <= k <= n2


Follow up:

	Could you solve the problem with a constant memory (i.e., O(1) memory complexity)?
	Could you solve the problem in O(n) time complexity? The solution may be too advanced for an interview
		but you may find reading this paper fun.
*/

// Define a type for the heap element
type Element struct {
	val int
	row int
	col int
}

// Define a PriorityQueue (min-heap) of Elements
type PriorityQueue []*Element

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].val < pq[j].val
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(*Element))
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	elem := old[n-1]
	*pq = old[0 : n-1]
	return elem
}

func kthSmallest(matrix [][]int, k int) int {
	n := len(matrix)
	pq := &PriorityQueue{}
	heap.Init(pq)

	// Initialize the heap with the first element of each row
	for i := 0; i < n; i++ {
		heap.Push(pq, &Element{val: matrix[i][0], row: i, col: 0})
	}

	// Extract the smallest element from the heap k times
	for i := 0; i < k-1; i++ {
		elem := heap.Pop(pq).(*Element)
		if elem.col < n-1 {
			heap.Push(pq, &Element{val: matrix[elem.row][elem.col+1], row: elem.row, col: elem.col + 1})
		}
	}

	// The root of the heap is the kth smallest element
	return heap.Pop(pq).(*Element).val
}

func main() {
	// Example 1
	matrix1 := [][]int{
		{1, 5, 9},
		{10, 11, 13},
		{12, 13, 15},
	}
	k1 := 8
	fmt.Println(kthSmallest(matrix1, k1)) // Output: 13

	// Example 2
	matrix2 := [][]int{
		{-5},
	}
	k2 := 1
	fmt.Println(kthSmallest(matrix2, k2)) // Output: -5
}
