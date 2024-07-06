package main

import (
	"container/heap"
	"fmt"
	"sort"
)

/*
Given an array of meeting time intervals intervals where intervals[i] = [starti, endi], return the
minimum number of conference rooms required.


Example 1:

	Input: intervals = [[0,30],[5,10],[15,20]]
	Output: 2

Example 2:

	Input: intervals = [[7,10],[2,4]]
	Output: 1


Constraints:

	1 <= intervals.length <= 10^4
	0 <= starti < endi <= 10^6

*/

// MeetingInterval represents a meeting with a start and end time.
type MeetingInterval struct {
	Start int
	End   int
}

// MinHeap is a min-heap of integers (end times).
type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// minMeetingRooms returns the minimum number of conference rooms required.
func minMeetingRooms(intervals [][]int) int {
	if len(intervals) == 0 {
		return 0
	}

	// Step 1: Separate start and end times
	startTimes := make([]int, len(intervals))
	endTimes := make([]int, len(intervals))
	for i, interval := range intervals {
		startTimes[i] = interval[0]
		endTimes[i] = interval[1]
	}

	// Step 2: Sort start and end times
	sort.Ints(startTimes)
	sort.Ints(endTimes)

	// Step 3: Use a min-heap to track the earliest end time
	minHeap := &MinHeap{}
	heap.Init(minHeap)
	heap.Push(minHeap, endTimes[0])

	roomCount := 1

	// Step 4: Iterate through the start times and manage rooms
	for i := 1; i < len(startTimes); i++ {
		// If the current meeting starts after the earliest ending meeting
		if startTimes[i] >= (*minHeap)[0] {
			heap.Pop(minHeap)
		} else {
			roomCount++
		}
		heap.Push(minHeap, endTimes[i])
	}

	return roomCount
}

func main() {
	// Example 1
	intervals1 := [][]int{{0, 30}, {5, 10}, {15, 20}}
	fmt.Println("Output:", minMeetingRooms(intervals1)) // Output: 2

	// Example 2
	intervals2 := [][]int{{7, 10}, {2, 4}}
	fmt.Println("Output:", minMeetingRooms(intervals2)) // Output: 1
}
