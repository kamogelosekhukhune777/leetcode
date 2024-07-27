package main

import "fmt"

/*
Given an integer array nums of unique elements, return all possible subsets (the power set).

The solution set must not contain duplicate subsets. Return the solution in any order.


Example 1:

	Input: nums = [1,2,3]
	Output: [[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]

Example 2:

	Input: nums = [0]
	Output: [[],[0]]


Constraints:

	1 <= nums.length <= 10
	-10 <= nums[i] <= 10
	All the numbers of nums are unique.
*/

// Function to generate all subsets
func subsets(nums []int) [][]int {
	var result [][]int // Initialize the result list
	var path []int     // Temporary list to store the current subset

	// Backtracking function to generate subsets
	var backtrack func(start int)
	backtrack = func(start int) {
		// Add a copy of the current subset to the result
		temp := make([]int, len(path))
		copy(temp, path)
		result = append(result, temp)

		// Iterate through the elements starting from the current index
		for i := start; i < len(nums); i++ {
			// Add the current element to the subset
			path = append(path, nums[i])
			// Recursively generate subsets with the current element
			backtrack(i + 1)
			// Remove the current element to backtrack
			path = path[:len(path)-1]
		}
	}

	// Start the backtracking process from the first element
	backtrack(0)
	return result
}

// Main function to test the subsets function
func main() {
	nums1 := []int{1, 2, 3}
	fmt.Println(subsets(nums1)) // Output: [[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]

	nums2 := []int{0}
	fmt.Println(subsets(nums2)) // Output: [[],[0]]
}
