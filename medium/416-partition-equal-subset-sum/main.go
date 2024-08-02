package main

import (
	"fmt"
)

/*
Given an integer array nums, return true if you can partition the array into two subsets such that
the sum of the elements in both subsets is equal or false otherwise.


Example 1:

	Input: nums = [1,5,11,5]
	Output: true
	Explanation: The array can be partitioned as [1, 5, 5] and [11].

Example 2:

	Input: nums = [1,2,3,5]
	Output: false
	Explanation: The array cannot be partitioned into equal sum subsets.


Constraints:

	1 <= nums.length <= 200
	1 <= nums[i] <= 100

*/
// Function to determine if the array can be partitioned into two subsets with equal sum
func canPartition(nums []int) bool {
	totalSum := 0
	for _, num := range nums {
		totalSum += num
	}

	// If total sum is odd, it cannot be partitioned into two equal subsets
	if totalSum%2 != 0 {
		return false
	}

	target := totalSum / 2
	dp := make([]bool, target+1)
	dp[0] = true

	// Fill the DP array
	for _, num := range nums {
		for j := target; j >= num; j-- {
			dp[j] = dp[j] || dp[j-num]
		}
	}

	return dp[target]
}

func main() {
	// Example 1
	nums1 := []int{1, 5, 11, 5}
	fmt.Println(canPartition(nums1)) // Output: true

	// Example 2
	nums2 := []int{1, 2, 3, 5}
	fmt.Println(canPartition(nums2)) // Output: false
}
