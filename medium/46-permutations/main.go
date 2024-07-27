package main

import "fmt"

/*
Given an array nums of distinct integers, return all the possible permutations. You
can return the answer in any order.

Example 1:

	Input: nums = [1,2,3]
	Output: [[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
Example 2:

	Input: nums = [0,1]
	Output: [[0,1],[1,0]]

Example 3:

	Input: nums = [1]
	Output: [[1]]


Constraints:

	1 <= nums.length <= 6
	-10 <= nums[i] <= 10
	All the integers of nums are unique.
*/

// Function to generate all permutations
func permute(nums []int) [][]int {
	var result [][]int // Initialize the result list

	// Backtracking function to generate permutations
	var backtrack func(start int)
	backtrack = func(start int) {
		// Base case: if start index reaches the length of nums, add a copy of nums to result
		if start == len(nums) {
			temp := make([]int, len(nums))
			copy(temp, nums)
			result = append(result, temp)
			return
		}

		// Iterate through the elements starting from the current index
		for i := start; i < len(nums); i++ {
			// Swap the current element with the starting element
			nums[start], nums[i] = nums[i], nums[start]
			// Recursively generate permutations with the current element
			backtrack(start + 1)
			// Swap back to restore the original array
			nums[start], nums[i] = nums[i], nums[start]
		}
	}

	// Start the backtracking process from the first element
	backtrack(0)
	return result
}

// Main function to test the permute function
func main() {
	nums1 := []int{1, 2, 3}
	fmt.Println(permute(nums1)) // Output: [[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]

	nums2 := []int{0, 1}
	fmt.Println(permute(nums2)) // Output: [[0,1],[1,0]]

	nums3 := []int{1}
	fmt.Println(permute(nums3)) // Output: [[1]]
}
