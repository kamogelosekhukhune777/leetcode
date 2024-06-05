package main

import "fmt"

/*
Given an array of integers nums containing n + 1 integers where each integer is in the range [1, n] inclusive.

There is only one repeated number in nums, return this repeated number.

You must solve the problem without modifying the array nums and uses only constant extra space.


Example 1:
	Input: nums = [1,3,4,2,2]
	Output: 2

Example 2:
	Input: nums = [3,1,3,4,2]
	Output: 3

Example 3:
	Input: nums = [3,3,3,3,3]
	Output: 3


Constraints:
	1 <= n <= 105
	nums.length == n + 1
	1 <= nums[i] <= n
	All the integers in nums appear only once except for precisely one
	 integer which appears two or more times.


Follow up:
	How can we prove that at least one duplicate number must exist in nums?
	Can you solve the problem in linear runtime complexity?
*/

func findDuplicate(nums []int) int {
	// Initialize slow and fast pointers
	slow := nums[0]
	fast := nums[0]

	// Phase 1: Finding the intersection point
	for {
		slow = nums[slow]
		fast = nums[nums[fast]]
		if slow == fast {
			break
		}
	}

	// Phase 2: Finding the entrance to the cycle
	slow = nums[0]
	for slow != fast {
		slow = nums[slow]
		fast = nums[fast]
	}

	return slow
}

func main() {
	// Test cases
	fmt.Println(findDuplicate([]int{1, 3, 4, 2, 2})) // Output: 2
	fmt.Println(findDuplicate([]int{3, 1, 3, 4, 2})) // Output: 3
	fmt.Println(findDuplicate([]int{3, 3, 3, 3, 3})) // Output: 3
}
