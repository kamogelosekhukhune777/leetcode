package main

import (
	"fmt"
	"sort"
)

/*

Given an array of intervals where intervals[i] = [starti, endi], merge
all overlapping intervals, and return an array of the non-overlapping intervals
that cover all the intervals in the input.


Example 1:
	Input: intervals = [[1,3],[2,6],[8,10],[15,18]]
	Output: [[1,6],[8,10],[15,18]]
	Explanation: Since intervals [1,3] and [2,6] overlap, merge them into [1,6].

Example 2:
	Input: intervals = [[1,4],[4,5]]
	Output: [[1,5]]
	Explanation: Intervals [1,4] and [4,5] are considered overlapping.


Constraints:

	1 <= intervals.length <= 104
	intervals[i].length == 2
	0 <= starti <= endi <= 104
*/

func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return [][]int{}
	}

	// Sort intervals based on the starting value
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	// Initialize the result list with the first interval
	merged := [][]int{intervals[0]}

	for _, interval := range intervals[1:] {
		// Get the last interval in the result list
		lastInterval := merged[len(merged)-1]

		// Check if the current interval overlaps with the last interval
		if interval[0] <= lastInterval[1] {
			// Merge the intervals by updating the end value
			if interval[1] > lastInterval[1] {
				lastInterval[1] = interval[1]
			}
		} else {
			// Add the current interval to the result list
			merged = append(merged, interval)
		}
	}

	return merged
}

func main() {
	intervals1 := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	merged1 := merge(intervals1)
	fmt.Println(merged1) // Output: [[1, 6], [8, 10], [15, 18]]

	intervals2 := [][]int{{1, 4}, {4, 5}}
	merged2 := merge(intervals2)
	fmt.Println(merged2) // Output: [[1, 5]]
}
