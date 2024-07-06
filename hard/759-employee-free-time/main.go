package main

import (
	"fmt"
	"sort"
)

/*
We are given a list schedule of employees, which represents the working time for each employee.

Each employee has a list of non-overlapping Intervals, and these intervals are in sorted order.

Return the list of finite intervals representing common, positive-length free time for all
employees, also in sorted order.

Example 1:

	Input: schedule = [[[1,2],[5,6]],[[1,3]],[[4,10]]]
	Output: [[3,4]]
	Explanation:
	There are a total of three employees, and all common
	free time intervals would be [-inf, 1], [3, 4], [10, inf].
	We discard any intervals that contain inf as they aren't finite.

Example 2:

	Input: schedule = [[[1,3],[6,7]],[[2,4]],[[2,5],[9,12]]]
	Output: [[5,6],[7,9]]

(Even though we are representing Intervals in the form [x, y], the objects inside are Intervals, not
lists or arrays. For example, schedule[0][0].start = 1, schedule[0][0].end = 2, and schedule[0][0][0] is not defined.)

Also, we wouldn't include intervals like [5, 5] in our answer, as they have zero length.

Note:

	schedule and schedule[i] are lists with lengths in range [1, 50].
	0 <= schedule[i].start < schedule[i].end <= 10^8.

*/

// Interval represents a working time interval with start and end times.
type Interval struct {
	Start, End int
}

// employeeFreeTime finds the common free time intervals for all employees.
func employeeFreeTime(schedule [][]Interval) []Interval {
	// Step 1: Flatten the schedule
	var intervals []Interval
	for _, employee := range schedule {
		intervals = append(intervals, employee...)
	}

	// Step 2: Sort the intervals by start times
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i].Start == intervals[j].Start {
			return intervals[i].End < intervals[j].End
		}
		return intervals[i].Start < intervals[j].Start
	})

	// Step 3: Merge overlapping intervals
	var merged []Interval
	for _, interval := range intervals {
		if len(merged) == 0 || merged[len(merged)-1].End < interval.Start {
			merged = append(merged, interval)
		} else {
			merged[len(merged)-1].End = max(merged[len(merged)-1].End, interval.End)
		}
	}

	// Step 4: Find the gaps between the merged intervals (free time)
	var freeTime []Interval
	for i := 1; i < len(merged); i++ {
		start := merged[i-1].End
		end := merged[i].Start
		if start < end {
			freeTime = append(freeTime, Interval{Start: start, End: end})
		}
	}

	return freeTime
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
	schedule1 := [][]Interval{
		{{1, 2}, {5, 6}},
		{{1, 3}},
		{{4, 10}},
	}
	result1 := employeeFreeTime(schedule1)
	fmt.Println("Output:", result1) // Output: [[3, 4]]

	// Example 2
	schedule2 := [][]Interval{
		{{1, 3}, {6, 7}},
		{{2, 4}},
		{{2, 5}, {9, 12}},
	}
	result2 := employeeFreeTime(schedule2)
	fmt.Println("Output:", result2) // Output: [[5, 6], [7, 9]]
}
