package main

import (
	"fmt"
)

/*
There is an integer array nums sorted in ascending order (with distinct values).

Prior to being passed to your function, nums is possibly rotated at an unknown pivot
index k (1 <= k < nums.length) such that the resulting array is [nums[k], nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]] (0-indexed). For example, [0,1,2,4,5,6,7] might be rotated at pivot index 3 and become [4,5,6,7,0,1,2].

Given the array nums after the possible rotation and an integer target, return the index of target
if it is in nums, or -1 if it is not in nums.

You must write an algorithm with O(log n) runtime complexity.


Example 1:

	Input: nums = [4,5,6,7,0,1,2], target = 0
	Output: 4
Example 2:

	Input: nums = [4,5,6,7,0,1,2], target = 3
	Output: -1
Example 3:

	Input: nums = [1], target = 0
	Output: -1

Constraints:

	1 <= nums.length <= 5000
	-104 <= nums[i] <= 104
	All values of nums are unique.
	nums is an ascending array that is possibly rotated.
	-104 <= target <= 104

*/

// Function to search for a target element in a rotated sorted array using modified binary search
func search(nums []int, target int) int {
	low, high := 0, len(nums)-1

	for low <= high {
		mid := low + (high-low)/2

		if nums[mid] == target {
			return mid // Target found, return its index
		}

		// Determine which part is sorted
		if nums[low] <= nums[mid] { // Left part is sorted
			if nums[low] <= target && target < nums[mid] {
				high = mid - 1 // Target is in the left part
			} else {
				low = mid + 1 // Target is in the right part
			}
		} else { // Right part is sorted
			if nums[mid] < target && target <= nums[high] {
				low = mid + 1 // Target is in the right part
			} else {
				high = mid - 1 // Target is in the left part
			}
		}
	}

	return -1 // Target not found
}

func main() {
	nums1 := []int{4, 5, 6, 7, 0, 1, 2}
	target1 := 0
	fmt.Printf("Index of %d in array: %d\n", target1, search(nums1, target1)) // Output: 4

	nums2 := []int{4, 5, 6, 7, 0, 1, 2}
	target2 := 3
	fmt.Printf("Index of %d in array: %d\n", target2, search(nums2, target2)) // Output: -1

	nums3 := []int{1}
	target3 := 0
	fmt.Printf("Index of %d in array: %d\n", target3, search(nums3, target3)) // Output: -1
}
