package main

import (
	"fmt"
)

//https://leetcode.com/problems/jump-game-ii/description/
/*
You are given a 0-indexed array of integers nums of length n. You
are initially positioned at nums[0].

Each element nums[i] represents the maximum length of a forward jump
from index i. In other words, if you are at nums[i], you can jump to any nums[i + j] where:

-	0 <= j <= nums[i] and
-	i + j < n
Return the minimum number of jumps to reach nums[n - 1]. The test cases are generated
such that you can reach nums[n - 1].



Example 1:

	Input: nums = [2,3,1,1,4]
	Output: 2
	Explanation: The minimum number of jumps to reach the last index is 2. Jump 1 step
		from index 0 to 1, then 3 steps to the last index.

Example 2:

	Input: nums = [2,3,0,1,4]
	Output: 2


Constraints:

	1 <= nums.length <= 10^4
	0 <= nums[i] <= 1000
	It's guaranteed that you can reach nums[n - 1].
*/

func jump(nums []int) int {
	n := len(nums)
	if n == 1 {
		return 0
	}

	jumps := 0
	currentEnd := 0
	farthest := 0

	for i := 0; i < n-1; i++ {
		// Update the farthest we can reach
		farthest = max(farthest, i+nums[i])

		// If we have reached the end of the current jump range
		if i == currentEnd {
			jumps++
			currentEnd = farthest
		}

		// If we can already reach the last index
		if currentEnd >= n-1 {
			break
		}
	}

	return jumps
}

// Helper function to get the maximum of two numbers
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	nums1 := []int{2, 3, 1, 1, 4}
	fmt.Println(jump(nums1)) // Output: 2

	nums2 := []int{2, 3, 0, 1, 4}
	fmt.Println(jump(nums2)) // Output: 2
}
