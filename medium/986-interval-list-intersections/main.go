package main

import (
	"fmt"
)

/*
You are given two lists of closed intervals, firstList and secondList,
where firstList[i] = [starti, endi] and secondList[j] = [startj, endj]. Each list of intervals
is pairwise disjoint and in sorted order.

Return the intersection of these two interval lists.

A closed interval [a, b] (with a <= b) denotes the set of real numbers x with a <= x <= b.

The intersection of two closed intervals is a set of real numbers that are either empty or represented
as a closed interval. For example, the intersection of [1, 3] and [2, 4] is [2, 3].


Example 1:

	Input: firstList = [[0,2],[5,10],[13,23],[24,25]], secondList = [[1,5],[8,12],[15,24],[25,26]]
	Output: [[1,2],[5,5],[8,10],[15,23],[24,24],[25,25]]

Example 2:

	Input: firstList = [[1,3],[5,9]], secondList = []
	Output: []

Constraints:

	0 <= firstList.length, secondList.length <= 1000
	firstList.length + secondList.length >= 1
	0 <= starti < endi <= 109
	endi < starti+1
	0 <= startj < endj <= 109
	endj < startj+1

*/

// intervalIntersection finds the intersections between two lists of intervals.
func intervalIntersection(firstList [][]int, secondList [][]int) [][]int {
	var result [][]int
	i, j := 0, 0
	n, m := len(firstList), len(secondList)

	// Step 1: Iterate through both lists
	for i < n && j < m {
		// Get the start and end points of the current intervals
		start1, end1 := firstList[i][0], firstList[i][1]
		start2, end2 := secondList[j][0], secondList[j][1]

		// Step 2: Find the intersection, if any
		start := max(start1, start2)
		end := min(end1, end2)
		if start <= end {
			// If there is an intersection, add it to the result
			result = append(result, []int{start, end})
		}

		// Step 3: Move the pointer with the smaller end time
		if end1 < end2 {
			i++
		} else {
			j++
		}
	}

	return result
}

// Helper function to find the minimum of two integers
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Helper function to find the maximum of two integers
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	// Example 1
	firstList1 := [][]int{{0, 2}, {5, 10}, {13, 23}, {24, 25}}
	secondList1 := [][]int{{1, 5}, {8, 12}, {15, 24}, {25, 26}}
	result1 := intervalIntersection(firstList1, secondList1)
	fmt.Println("Output:", result1) // Output: [[1, 2], [5, 5], [8, 10], [15, 23], [24, 24], [25, 25]]

	// Example 2
	firstList2 := [][]int{{1, 3}, {5, 9}}
	secondList2 := [][]int{}
	result2 := intervalIntersection(firstList2, secondList2)
	fmt.Println("Output:", result2) // Output: []
}
