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

// PriorityQueue implements heap.Interface and holds Pairs
type PriorityQueue []Pair

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].sum < pq[j].sum
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(Pair))
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	if len(nums1) == 0 || len(nums2) == 0 || k == 0 {
		return [][]int{}
	}

	pq := &PriorityQueue{}
	heap.Init(pq)

	// Initialize the heap with pairs consisting of the first element of nums1 and every element in nums2
	for j := 0; j < len(nums2) && j < k; j++ {
		heap.Push(pq, Pair{sum: nums1[0] + nums2[j], i: 0, j: j})
	}

	result := [][]int{}

	// Extract the smallest pairs from the heap up to k times
	for len(result) < k && pq.Len() > 0 {
		p := heap.Pop(pq).(Pair)
		result = append(result, []int{nums1[p.i], nums2[p.j]})

		// If possible, push the next pair from nums1 with the same element from nums2
		if p.i+1 < len(nums1) {
			heap.Push(pq, Pair{sum: nums1[p.i+1] + nums2[p.j], i: p.i + 1, j: p.j})
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
