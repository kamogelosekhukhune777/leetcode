package main

import (
	"fmt"
)

/*
Given an array of distinct integers candidates and a target integer target, return a list of all unique combinations
of candidates where the chosen numbers sum to target. You may return the combinations in any order.

The same number may be chosen from candidates an unlimited number of times. Two combinations are unique if the

frequency of at least one of the chosen numbers is different.

The test cases are generated such that the number of unique combinations that sum up to target is less than 150
combinations for the given input.


Example 1:

	Input: candidates = [2,3,6,7], target = 7
	Output: [[2,2,3],[7]]
	Explanation:
		2 and 3 are candidates, and 2 + 2 + 3 = 7. Note that 2 can be used multiple times.
		7 is a candidate, and 7 = 7.
		These are the only two combinations.

Example 2:

	Input: candidates = [2,3,5], target = 8
	Output: [[2,2,2,2],[2,3,3],[3,5]]

Example 3:

	Input: candidates = [2], target = 1
	Output: []


Constraints:

	1 <= candidates.length <= 30
	2 <= candidates[i] <= 40
	All elements of candidates are distinct.
	1 <= target <= 40
*/

// Function to find the subarray with the largest product
func maxProduct(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	// Initialize maxProd, minProd, and result with the first element
	maxProd, minProd, result := nums[0], nums[0], nums[0]

	// Iterate through the array starting from the second element
	for i := 1; i < len(nums); i++ {
		// If current element is negative, swap maxProd and minProd
		if nums[i] < 0 {
			maxProd, minProd = minProd, maxProd
		}

		// Update maxProd and minProd
		maxProd = max(nums[i], maxProd*nums[i])
		minProd = min(nums[i], minProd*nums[i])

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
