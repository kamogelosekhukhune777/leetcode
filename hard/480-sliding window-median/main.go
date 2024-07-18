package main

import (
	"container/heap"
	"fmt"
)

/*
The median is the middle value in an ordered integer list. If the size of the list is even, there is no middle value.
So the median is the mean of the two middle values.

For examples, if arr = [2,3,4], the median is 3.
For examples, if arr = [1,2,3,4], the median is (2 + 3) / 2 = 2.5.
You are given an integer array nums and an integer k. There is a sliding window of size k which is moving from the very
left of the array to the very right. You can only see the k numbers in the window. Each time the sliding window moves
right by one position.

Return the median array for each window in the original array. Answers within 10-5 of the actual value will be accepted.


Example 1:

	Input: nums = [1,3,-1,-3,5,3,6,7], k = 3
	Output: [1.00000,-1.00000,-1.00000,3.00000,5.00000,6.00000]
	Explanation:
		Window position                Median
		---------------                -----
		[1  3  -1] -3  5  3  6  7        1
		1 [3  -1  -3] 5  3  6  7       -1
		1  3 [-1  -3  5] 3  6  7       -1
		1  3  -1 [-3  5  3] 6  7        3
		1  3  -1  -3 [5  3  6] 7        5
		1  3  -1  -3  5 [3  6  7]       6

Example 2:

	Input: nums = [1,2,3,4,2,3,1,4,2], k = 3
	Output: [2.00000,3.00000,3.00000,3.00000,2.00000,3.00000,2.00000]


Constraints:

	1 <= k <= nums.length <= 10^5
	-2^31 <= nums[i] <= 2^31 - 1

*/

type MaxHeap []int

func (h MaxHeap) Len() int {
	return len(h)
}
func (h MaxHeap) Less(i, j int) bool {
	return h[i] > h[j]
}
func (h MaxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type MinHeap []int

func (h MinHeap) Len() int {
	return len(h)
}

func (h MinHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

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

type MedianFinder struct {
	maxHeap *MaxHeap
	minHeap *MinHeap
}

func Constructor() MedianFinder {
	maxHeap := &MaxHeap{}
	minHeap := &MinHeap{}
	heap.Init(maxHeap)
	heap.Init(minHeap)
	return MedianFinder{maxHeap, minHeap}
}

func (mf *MedianFinder) AddNum(num int) {
	heap.Push(mf.maxHeap, num)
	heap.Push(mf.minHeap, heap.Pop(mf.maxHeap))
	if mf.maxHeap.Len() < mf.minHeap.Len() {
		heap.Push(mf.maxHeap, heap.Pop(mf.minHeap))
	}
}

func (mf *MedianFinder) RemoveNum(num int) {
	if num <= (*mf.maxHeap)[0] {
		for i := 0; i < mf.maxHeap.Len(); i++ {
			if (*mf.maxHeap)[i] == num {
				heap.Remove(mf.maxHeap, i)
				break
			}
		}
	} else {
		for i := 0; i < mf.minHeap.Len(); i++ {
			if (*mf.minHeap)[i] == num {
				heap.Remove(mf.minHeap, i)
				break
			}
		}
	}
	if mf.maxHeap.Len() < mf.minHeap.Len() {
		heap.Push(mf.maxHeap, heap.Pop(mf.minHeap))
	} else if mf.maxHeap.Len() > mf.minHeap.Len()+1 {
		heap.Push(mf.minHeap, heap.Pop(mf.maxHeap))
	}
}

func (mf *MedianFinder) FindMedian() float64 {
	if mf.maxHeap.Len() > mf.minHeap.Len() {
		return float64((*mf.maxHeap)[0])
	}
	return float64((*mf.maxHeap)[0]+(*mf.minHeap)[0]) / 2.0
}

func medianSlidingWindow(nums []int, k int) []float64 {
	mf := Constructor()
	result := []float64{}
	for i := 0; i < k; i++ {
		mf.AddNum(nums[i])
	}
	result = append(result, mf.FindMedian())
	for i := k; i < len(nums); i++ {
		mf.RemoveNum(nums[i-k])
		mf.AddNum(nums[i])
		result = append(result, mf.FindMedian())
	}
	return result
}

func main() {
	nums := []int{1, 3, -1, -3, 5, 3, 6, 7}
	k := 3
	fmt.Println(medianSlidingWindow(nums, k)) // Output: [1.00000,-1.00000,-1.00000,3.00000,5.00000,6.00000]
}
