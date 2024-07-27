package main

import (
	"fmt"
	"sort"
)

//
/*
Given an integer array nums and an integer k, return true if it is possible
to divide this array into k non-empty subsets whose sums are all equal.

Example 1:

	Input: nums = [4,3,2,3,5,2,1], k = 4
	Output: true
	Explanation: It is possible to divide it into 4 subsets (5), (1, 4), (2,3), (2,3) with equal sums.

Example 2:

	Input: nums = [1,2,3,4], k = 3
	Output: false


Constraints:

	1 <= k <= nums.length <= 16
	1 <= nums[i] <= 104
	The frequency of each element is in the range [1, 4].

*/

// Function to determine if the array can be partitioned into k subsets with equal sums
func canPartitionKSubsets(nums []int, k int) bool {
	totalSum := 0
	for _, num := range nums {
		totalSum += num
	}

	// If total sum is not divisible by k, return false
	if totalSum%k != 0 {
		return false
	}

	target := totalSum / k                       // Each subset must sum to target
	sort.Sort(sort.Reverse(sort.IntSlice(nums))) // Sort the array in descending order

	// Initialize an array to keep track of the sum of each subset
	subsets := make([]int, k)

	// Backtracking function to try placing each number in a subset
	var backtrack func(index int) bool
	backtrack = func(index int) bool {
		// Base case: if all numbers are placed
		if index == len(nums) {
			for _, subsetSum := range subsets {
				if subsetSum != target {
					return false
				}
			}
			return true
		}

		// Try placing the current number in each subset
		for i := 0; i < k; i++ {
			if subsets[i]+nums[index] <= target {
				subsets[i] += nums[index]
				if backtrack(index + 1) {
					return true
				}
				subsets[i] -= nums[index]
			}
			// If this subset is empty, don't try further (optimization)
			if subsets[i] == 0 {
				break
			}
		}
		return false
	}

	return backtrack(0)
}

// Main function to test the canPartitionKSubsets function
func main() {
	nums1 := []int{4, 3, 2, 3, 5, 2, 1}
	k1 := 4
	fmt.Println(canPartitionKSubsets(nums1, k1)) // Output: true

	nums2 := []int{1, 2, 3, 4}
	k2 := 3
	fmt.Println(canPartitionKSubsets(nums2, k2)) // Output: false
}
