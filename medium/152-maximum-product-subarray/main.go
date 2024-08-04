package main

import (
	"fmt"
)

/*
Given an integer array nums, find a
subarray that has the largest product, and return the product.
The test cases are generated so that the answer will fit in a 32-bit integer.

Example 1:

	Input: nums = [2,3,-2,4]
	Output: 6
	Explanation: [2,3] has the largest product 6.

Example 2:

	Input: nums = [-2,0,-1]
	Output: 0
	Explanation: The result cannot be 2, because [-2,-1] is not a subarray.

Constraints:

	1 <= nums.length <= 2 * 104
	-10 <= nums[i] <= 10
	The product of any prefix or suffix of nums is guaranteed to fit in a 32-bit integer.
*/

// Function to find the subarray with the largest product
func maxProduct(nums []int) int {
	// Handle the case where nums is empty (not required per problem constraints)
	if len(nums) == 0 {
		return 0
	}

	// Initialize maxProd, minProd, and result with the first element of the array
	maxProd, minProd, result := nums[0], nums[0], nums[0]

	// Iterate through the array starting from the second element
	for i := 1; i < len(nums); i++ {
		// Store current maxProd and minProd before updating them
		currentMaxProd := maxProd
		currentMinProd := minProd

		// Update maxProd and minProd
		maxProd = max(nums[i], max(currentMaxProd*nums[i], currentMinProd*nums[i]))
		minProd = min(nums[i], min(currentMaxProd*nums[i], currentMinProd*nums[i]))

		// Update result with the maximum product found so far
		result = max(result, maxProd)
	}

	return result
}

// Helper functions to get the maximum and minimum of two integers
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	// Example 1
	nums1 := []int{2, 3, -2, 4}
	fmt.Println(maxProduct(nums1)) // Output: 6

	// Example 2
	nums2 := []int{-2, 0, -1}
	fmt.Println(maxProduct(nums2)) // Output: 0

	// Additional Example
	nums3 := []int{-2, 3, -4}
	fmt.Println(maxProduct(nums3)) // Output: 24
}
