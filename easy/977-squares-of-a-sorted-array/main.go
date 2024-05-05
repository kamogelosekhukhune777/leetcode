package main

import "fmt"

/*
Given an integer array nums sorted in non-decreasing order, return an array of the squares
of each number sorted in non-decreasing order.

Example 1:
	Input: nums = [-4,-1,0,3,10]
	Output: [0,1,9,16,100]
	Explanation: After squaring, the array becomes [16,1,0,9,100].
	After sorting, it becomes [0,1,9,16,100].

Example 2:
	Input: nums = [-7,-3,2,3,11]
	Output: [4,9,9,49,121]

Constraints:
	1 <= nums.length <= 104
	-104 <= nums[i] <= 104
	nums is sorted in non-decreasing order.
*/
func sortedSquares(nums []int) []int {
	n := len(nums)
	result := make([]int, n)

	// Initialize pointers
	left, right := 0, n-1

	// Fill the result array from right to left with squared values
	for i := n - 1; i >= 0; i-- {
		if abs(nums[left]) > abs(nums[right]) {
			result[i] = nums[left] * nums[left]
			left++
		} else {
			result[i] = nums[right] * nums[right]
			right--
		}
	}

	return result
}

// Helper function to get absolute value
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	//Example 1:
	nums1 := []int{-4, -1, 0, 3, 10}
	fmt.Println(sortedSquares(nums1))

	//Example 2:
	nums := []int{-7, -3, 2, 3, 11}
	fmt.Println(sortedSquares(nums))
}
