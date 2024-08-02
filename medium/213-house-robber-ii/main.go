package main

import (
	"fmt"
)

/*
You are a professional robber planning to rob houses along a street. Each house has a certain
amount of money stashed. All houses at this place are arranged in a circle. That means the first
house is the neighbor of the last one. Meanwhile, adjacent houses have a security system
connected, and it will automatically contact the police if two adjacent houses were broken into
on the same night.

Given an integer array nums representing the amount of money of each house, return the maximum amount
of money you can rob tonight without alerting the police.


Example 1:

	Input: nums = [2,3,2]
	Output: 3
	Explanation: You cannot rob house 1 (money = 2) and then rob
		house 3 (money = 2), because they are adjacent houses.

Example 2:

	Input: nums = [1,2,3,1]
	Output: 4
	Explanation: Rob house 1 (money = 1) and then rob house 3 (money = 3).
		Total amount you can rob = 1 + 3 = 4.
Example 3:

	Input: nums = [1,2,3]
	Output: 3


Constraints:

	1 <= nums.length <= 100
	0 <= nums[i] <= 1000
*/

// Helper function to compute the maximum amount of money that can be robbed in a linear sequence of houses
func robLinear(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return nums[0]
	}
	dp := make([]int, n)
	dp[0] = nums[0]
	dp[1] = max(nums[0], nums[1])
	for i := 2; i < n; i++ {
		dp[i] = max(dp[i-1], dp[i-2]+nums[i])
	}
	return dp[n-1]
}

// Function to compute the maximum amount of money that can be robbed in a circular arrangement of houses
func rob(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return nums[0]
	}
	// Case 1: Rob houses from the first to the second last house
	max1 := robLinear(nums[:n-1])
	// Case 2: Rob houses from the second to the last house
	max2 := robLinear(nums[1:])
	return max(max1, max2)
}

// Helper function to get the maximum of two integers
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	// Example 1
	nums1 := []int{2, 3, 2}
	fmt.Println(rob(nums1)) // Output: 3

	// Example 2
	nums2 := []int{1, 2, 3, 1}
	fmt.Println(rob(nums2)) // Output: 4

	// Example 3
	nums3 := []int{1, 2, 3}
	fmt.Println(rob(nums3)) // Output: 3
}
