package main

import (
	"fmt"
)

/*
There is an integer array nums sorted in non-decreasing order (not necessarily with distinct values).

Before being passed to your function, nums is rotated at an unknown pivot index k (0 <= k < nums.length)
such that the resulting array is [nums[k], nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]] (0-indexed).
For example, [0,1,2,4,4,4,5,6,6,7] might be rotated at pivot index 5 and become [4,5,6,6,7,0,1,2,4,4].

Given the array nums after the rotation and an integer target, return true if target is in nums, or false if
it is not in nums.

You must decrease the overall operation steps as much as possible.


Example 1:

	Input: nums = [2,5,6,0,0,1,2], target = 0
	Output: true

Example 2:

	Input: nums = [2,5,6,0,0,1,2], target = 3
	Output: false


Constraints:

	1 <= nums.length <= 5000
	-10^4 <= nums[i] <= 10^4
	nums is guaranteed to be rotated at some pivot.
	-10^4 <= target <= 10^4


Follow up: This problem is similar to Search in Rotated Sorted Array, but nums may contain
	duplicates. Would this affect the runtime complexity? How and why?
*/

// Function to search for a target in a rotated sorted array with duplicates
func search(nums []int, target int) bool {
	low, high := 0, len(nums)-1

	for low <= high {
		mid := low + (high-low)/2

		if nums[mid] == target {
			return true
		}

		// If we can't determine the sorted half, adjust the pointers
		if nums[low] == nums[mid] && nums[mid] == nums[high] {
			low++
			high--
		} else if nums[low] <= nums[mid] { // Left half is sorted
			if nums[low] <= target && target < nums[mid] {
				high = mid - 1
			} else {
				low = mid + 1
			}
		} else { // Right half is sorted
			if nums[mid] < target && target <= nums[high] {
				low = mid + 1
			} else {
				high = mid - 1
			}
		}
	}

	return false
}

func main() {
	nums1 := []int{2, 5, 6, 0, 0, 1, 2}
	target1 := 0
	fmt.Println(search(nums1, target1)) // Output: true

	nums2 := []int{2, 5, 6, 0, 0, 1, 2}
	target2 := 3
	fmt.Println(search(nums2, target2)) // Output: false
}
