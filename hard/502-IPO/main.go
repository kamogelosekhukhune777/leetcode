package main

import (
	"container/heap"
	"fmt"
)

/*
Suppose LeetCode will start its IPO soon. In order to sell a good price of its shares to Venture Capital, LeetCode
would like to work on some projects to increase its capital before the IPO. Since it has limited resources, it can only finish
at most k distinct projects before the IPO. Help LeetCode design the best way to maximize its total capital after finishing at
most k distinct projects.

You are given n projects where the ith project has a pure profit profits[i] and a minimum capital of capital[i] is needed to start it.

Initially, you have w capital. When you finish a project, you will obtain its pure profit and the profit will be added
to your total capital.

Pick a list of at most k distinct projects from given projects to maximize your final capital, and
return the final maximized capital.

The answer is guaranteed to fit in a 32-bit signed integer.


Example 1:

	Input: k = 2, w = 0, profits = [1,2,3], capital = [0,1,1]
	Output: 4
	Explanation: Since your initial capital is 0, you can only start the project indexed 0.
		After finishing it you will obtain profit 1 and your capital becomes 1.
		With capital 1, you can either start the project indexed 1 or the project indexed 2.
		Since you can choose at most 2 projects, you need to finish the project indexed 2 to get the maximum capital.
		Therefore, output the final maximized capital, which is 0 + 1 + 3 = 4.

Example 2:

	Input: k = 3, w = 0, profits = [1,2,3], capital = [0,1,2]
	Output: 6


Constraints:

	1 <= k <= 105
	0 <= w <= 109
	n == profits.length
	n == capital.length
	1 <= n <= 105
	0 <= profits[i] <= 104
	0 <= capital[i] <= 109

*/

// MinHeap for capital
type MinHeap [][]int

func (h MinHeap) Len() int {
	return len(h)
}

func (h MinHeap) Less(i, j int) bool {
	return h[i][0] < h[j][0]
}

func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.([]int))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// MaxHeap for profits
type MaxHeap [][]int

func (h MaxHeap) Len() int {
	return len(h)
}

func (h MaxHeap) Less(i, j int) bool {
	return h[i][1] > h[j][1]
}

func (h MaxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.([]int))
}

func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func findMaximizedCapital(k int, w int, profits []int, capital []int) int {
	minHeap := &MinHeap{}
	maxHeap := &MaxHeap{}

	// Push all projects into the min-heap based on their capital requirement
	for i := 0; i < len(profits); i++ {
		heap.Push(minHeap, []int{capital[i], profits[i]})
	}

	// Iterate for at most k projects
	for i := 0; i < k; i++ {
		// Transfer all affordable projects to the max-heap
		for minHeap.Len() > 0 && (*minHeap)[0][0] <= w {
			heap.Push(maxHeap, heap.Pop(minHeap))
		}

		// If no projects can be started, break
		if maxHeap.Len() == 0 {
			break
		}

		// Start the most profitable project
		w += heap.Pop(maxHeap).([]int)[1]
	}

	return w
}

func main() {
	k := 2
	w := 0
	profits := []int{1, 2, 3}
	capital := []int{0, 1, 1}
	fmt.Println(findMaximizedCapital(k, w, profits, capital)) // Output: 4
}
