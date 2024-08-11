package main

import (
	"container/heap"
	"fmt"
)

/*
You are given two integer arrays nums1 and nums2 sorted in non-decreasing order and an integer k.

Define a pair (u, v) which consists of one element from the first array and one element from the second array.

Return the k pairs (u1, v1), (u2, v2), ..., (uk, vk) with the smallest sums.



Example 1:

	Input: nums1 = [1,7,11], nums2 = [2,4,6], k = 3
	Output: [[1,2],[1,4],[1,6]]
	Explanation: The first 3 pairs are returned from the sequence: [1,2],[1,4],[1,6],[7,2],[7,4],[11,2],[7,6],[11,4],[11,6]
Example 2:

	Input: nums1 = [1,1,2], nums2 = [1,2,3], k = 2
	Output: [[1,1],[1,1]]
	Explanation: The first 2 pairs are returned from the sequence: [1,1],[1,1],[1,2],[2,1],[1,2],[2,2],[1,3],[1,3],[2,3]


Constraints:

	1 <= nums1.length, nums2.length <= 105
	-109 <= nums1[i], nums2[i] <= 109
	nums1 and nums2 both are sorted in non-decreasing order.
	1 <= k <= 104
	k <= nums1.length * nums2.length
*/

// Pair struct to store the sum and indices of elements from nums1 and nums2
type Pair struct {
	sum int
	i   int // index in nums1
	j   int // index in nums2
}
type MinHeap []Pair

func (h MinHeap) Len() int {
	return len(h)
}

func (h MinHeap) Less(i, j int) bool {
	return h[i].sum < h[j].sum
}

func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(Pair))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	if len(nums1) == 0 || len(nums2) == 0 || k == 0 {
		return [][]int{}
	}

	minHeap := &MinHeap{}
	heap.Init(minHeap)

	// Start by pushing the first element of nums1 with each element of nums2
	for j := 0; j < len(nums2) && j < k; j++ {
		heap.Push(minHeap, Pair{nums1[0] + nums2[j], 0, j})
	}

	result := [][]int{}

	for k > 0 && minHeap.Len() > 0 {
		k--
		smallest := heap.Pop(minHeap).(Pair)
		i, j := smallest.i, smallest.j
		result = append(result, []int{nums1[i], nums2[j]})

		// If possible, push the next element in nums1 with nums2[j]
		if i+1 < len(nums1) {
			heap.Push(minHeap, Pair{nums1[i+1] + nums2[j], i + 1, j})
		}
	}

	return result
}

func main() {
	// Example 1
	nums1 := []int{1, 7, 11}
	nums2 := []int{2, 4, 6}
	k := 3
	fmt.Println(kSmallestPairs(nums1, nums2, k)) // Output: [[1, 2], [1, 4], [1, 6]]

	// Example 2
	nums1 = []int{1, 1, 2}
	nums2 = []int{1, 2, 3}
	k = 2
	fmt.Println(kSmallestPairs(nums1, nums2, k)) // Output: [[1, 1], [1, 1]]

	// Example 3
	nums1 = []int{1, 2}
	nums2 = []int{3}
	k = 3
	fmt.Println(kSmallestPairs(nums1, nums2, k)) // Output: [[1, 3], [2, 3]]
}
