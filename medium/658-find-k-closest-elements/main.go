package main

import (
	"fmt"
	"sort"
)

/*
Given a sorted integer array arr, two integers k and x, return the k closest integers to x
in the array. The result should also be sorted in ascending order.

An integer a is closer to x than an integer b if:

|a - x| < |b - x|, or
|a - x| == |b - x| and a < b


Example 1:

	Input: arr = [1,2,3,4,5], k = 4, x = 3
	Output: [1,2,3,4]

Example 2:

	Input: arr = [1,2,3,4,5], k = 4, x = -1
	Output: [1,2,3,4]


Constraints:

	1 <= k <= arr.length
	1 <= arr.length <= 10^4
	arr is sorted in ascending order.
	-10^4 <= arr[i], x <= 10^4
*/

// Function to find the k closest elements to x in arr
func findClosestElements(arr []int, k int, x int) []int {
	// Helper function to find the closest index using binary search
	idx := sort.Search(len(arr), func(i int) bool {
		return arr[i] >= x
	})

	// Initialize two pointers
	left, right := idx-1, idx

	// Collect k closest elements
	for i := 0; i < k; i++ {
		if left < 0 {
			right++
		} else if right >= len(arr) {
			left--
		} else if x-arr[left] <= arr[right]-x {
			left--
		} else {
			right++
		}
	}

	// Slice and return the result
	return arr[left+1 : right]
}

func main() {
	arr1 := []int{1, 2, 3, 4, 5}
	k1, x1 := 4, 3
	fmt.Println(findClosestElements(arr1, k1, x1)) // Output: [1, 2, 3, 4]

	arr2 := []int{1, 2, 3, 4, 5}
	k2, x2 := 4, -1
	fmt.Println(findClosestElements(arr2, k2, x2)) // Output: [1, 2, 3, 4]
}
