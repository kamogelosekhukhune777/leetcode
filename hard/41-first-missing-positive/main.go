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

	i := 0
	for i < n {
		correctIdx := nums[i] - 1
		// Only swap if nums[i] is in the range [1, n] and not already in the correct position
		if nums[i] > 0 && nums[i] <= n && nums[i] != nums[correctIdx] {
			nums[i], nums[correctIdx] = nums[correctIdx], nums[i]
		} else {
			i++
		}
	}

	// After the cyclic sort, find the first index where the index does not match nums[i]
	for i = 0; i < n; i++ {
		if nums[i] != i+1 {
			return i + 1
		}
	}

	// If all numbers from 1 to n are in their correct positions, the missing number is n+1
	return n + 1
}

func main() {
	// Test cases
	fmt.Println(firstMissingPositive([]int{1, 2, 0}))         // Output: 3
	fmt.Println(firstMissingPositive([]int{3, 4, -1, 1}))     // Output: 2
	fmt.Println(firstMissingPositive([]int{7, 8, 9, 11, 12})) // Output: 1
}
