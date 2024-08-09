package main

import "fmt"

/*
Given an unsorted integer array nums. Return the smallest positive integer that is not present in nums.

You must implement an algorithm that runs in O(n) time and uses O(1) auxiliary space.

Example 1:
	Input: nums = [1,2,0]
	Output: 3
	Explanation: The numbers in the range [1,2] are all in the array.

Example 2:
	Input: nums = [3,4,-1,1]
	Output: 2
	Explanation: 1 is in the array but 2 is missing.

Example 3:
	Input: nums = [7,8,9,11,12]
	Output: 1
	Explanation: The smallest positive integer 1 is missing.


Constraints:
	1 <= nums.length <= 105
	-231 <= nums[i] <= 231 - 1
*/

func firstMissingPositive(nums []int) int {
	n := len(nums)

	// Step 1: Place each number in its correct position if possible
	for i := 0; i < n; i++ {
		for nums[i] > 0 && nums[i] <= n && nums[nums[i]-1] != nums[i] {
			nums[nums[i]-1], nums[i] = nums[i], nums[nums[i]-1]
		}
	}

	// Step 2: Find the first missing positive integer
	for i := 0; i < n; i++ {
		if nums[i] != i+1 {
			return i + 1
		}
	}

	// Step 3: If all positions are correctly placed, the missing integer is n+1
	return n + 1
}

func main() {
	// Test cases
	fmt.Println(firstMissingPositive([]int{1, 2, 0}))         // Output: 3
	fmt.Println(firstMissingPositive([]int{3, 4, -1, 1}))     // Output: 2
	fmt.Println(firstMissingPositive([]int{7, 8, 9, 11, 12})) // Output: 1
}
