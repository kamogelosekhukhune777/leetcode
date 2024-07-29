package main

import "fmt"

/*
You are given an integer array nums. You are initially positioned at the array's
first index, and each element in the array represents your maximum
jump length at that position.

Return true if you can reach the last index, or false otherwise.

Example 1:

	Input: nums = [2,3,1,1,4]
	Output: true
	Explanation: Jump 1 step from index 0 to 1, then 3 steps to the last index.

Example 2:

	Input: nums = [3,2,1,0,4]
	Output: false
	Explanation: You will always arrive at index 3 no matter what. Its maximum
		jump length is 0, which makes it impossible to reach the last index.


Constraints:

	1 <= nums.length <= 104
	0 <= nums[i] <= 105
*/

// canJump determines if you can reach the last index from the first index
func canJump(nums []int) bool {
	// maxReach represents the farthest index we can reach
	maxReach := 0
	for i := 0; i < len(nums); i++ {
		// If the current index is greater than maxReach, we can't move further
		if i > maxReach {
			return false
		}
		// Update maxReach if the current index + nums[i] is farther
		if i+nums[i] > maxReach {
			maxReach = i + nums[i]
		}
		// If maxReach is already greater than or equal to the last index, return true
		if maxReach >= len(nums)-1 {
			return true
		}
	}
	return false
}

func main() {
	nums1 := []int{2, 3, 1, 1, 4}
	nums2 := []int{3, 2, 1, 0, 4}

	fmt.Println(canJump(nums1)) // Output: true
	fmt.Println(canJump(nums2)) // Output: false
}
