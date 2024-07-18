package main

import (
	"container/heap"
	"fmt"
)

/*
The median is the middle value in an ordered integer list. If the size of the list is even, there is no middle value, and the median is the mean of the two middle values.

For example, for arr = [2,3,4], the median is 3.
For example, for arr = [2,3], the median is (2 + 3) / 2 = 2.5.
Implement the MedianFinder class:

MedianFinder() initializes the MedianFinder object.
void addNum(int num) adds the integer num from the data stream to the data structure.
double findMedian() returns the median of all elements so far. Answers within 10-5 of the actual answer will be accepted.


Example 1:
	Input:
		["MedianFinder", "addNum", "addNum", "findMedian", "addNum", "findMedian"]
		[[], [1], [2], [], [3], []]
	Output:
		[null, null, null, 1.5, null, 2.0]

	Explanation:
		MedianFinder medianFinder = new MedianFinder();
		medianFinder.addNum(1);    // arr = [1]
		medianFinder.addNum(2);    // arr = [1, 2]
		medianFinder.findMedian(); // return 1.5 (i.e., (1 + 2) / 2)
		medianFinder.addNum(3);    // arr[1, 2, 3]
		medianFinder.findMedian(); // return 2.0

Constraints:
	-105 <= num <= 105
	There will be at least one element in the data structure before calling findMedian.
	At most 5 * 104 calls will be made to addNum and findMedian.

Follow up:
	If all integer numbers from the stream are in the range [0, 100], how would you optimize your solution?
	If 99% of all integer numbers from the stream are in the range [0, 100], how would you optimize your solution?

*/

// A max-heap for the lower half of numbers
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

// A min-heap for the upper half of numbers
type MinHeap []int

func (h MinHeap) Len() int {
	return len(h)
}

// Min-heap, so use <
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

// MedianFinder struct that uses two heaps
type MedianFinder struct {
	maxHeap *MaxHeap
	minHeap *MinHeap
}

// Constructor to initialize the MedianFinder
func Constructor() MedianFinder {
	maxHeap := &MaxHeap{}
	minHeap := &MinHeap{}
	heap.Init(maxHeap)
	heap.Init(minHeap)
	return MedianFinder{maxHeap, minHeap}
}

// AddNum adds a number to the data structure
func (mf *MedianFinder) AddNum(num int) {
	// Push the new number onto the max-heap
	heap.Push(mf.maxHeap, num)
	// Ensure the max-heap's max element is not greater than the min-heap's min element
	heap.Push(mf.minHeap, heap.Pop(mf.maxHeap))

	// Balance the heaps if necessary (max-heap can have at most one more element than min-heap)
	if mf.maxHeap.Len() < mf.minHeap.Len() {
		heap.Push(mf.maxHeap, heap.Pop(mf.minHeap))
	}
}

// FindMedian returns the median of all elements so far
func (mf *MedianFinder) FindMedian() float64 {
	if mf.maxHeap.Len() > mf.minHeap.Len() {
		return float64((*mf.maxHeap)[0])
	}
	return float64((*mf.maxHeap)[0]+(*mf.minHeap)[0]) / 2.0
}

func main() {
	medianFinder := Constructor()
	medianFinder.AddNum(1)                 // arr = [1]
	medianFinder.AddNum(2)                 // arr = [1, 2]
	fmt.Println(medianFinder.FindMedian()) // return 1.5
	medianFinder.AddNum(3)                 // arr = [1, 2, 3]
	fmt.Println(medianFinder.FindMedian()) // return 2.0
}
