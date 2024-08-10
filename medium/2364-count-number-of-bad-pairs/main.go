package main

import "fmt"

/*
You are given a 0-indexed integer array nums. A pair of indices (i, j) is
a bad pair if i < j and j - i != nums[j] - nums[i].

Return the total number of bad pairs in nums.


Example 1:

	Input: nums = [4,1,3,3]
	Output: 5
	Explanation: The pair (0, 1) is a bad pair since 1 - 0 != 1 - 4.
		The pair (0, 2) is a bad pair since 2 - 0 != 3 - 4, 2 != -1.
		The pair (0, 3) is a bad pair since 3 - 0 != 3 - 4, 3 != -1.
		The pair (1, 2) is a bad pair since 2 - 1 != 3 - 1, 1 != 2.
		The pair (2, 3) is a bad pair since 3 - 2 != 3 - 3, 1 != 0.
		There are a total of 5 bad pairs, so we return 5.

Example 2:

	Input: nums = [1,2,3,4,5]
	Output: 0
	Explanation: There are no bad pairs.


Constraints:

	1 <= nums.length <= 10^5
	1 <= nums[i] <= 10^9
*/

// Function to find the total number of bad pairs in the array
func countBadPairs(nums []int) int {
	n := len(nums)
	totalPairs := n * (n - 1) / 2

	countMap := make(map[int]int)
	goodPairs := 0

	for i, num := range nums {
		diff := num - i
		if val, exists := countMap[diff]; exists {
			goodPairs += val
		}
		countMap[diff]++
	}

	return totalPairs - goodPairs
}

func main() {
	// Test cases
	fmt.Println(countBadPairs([]int{4, 1, 3, 3}))    // Output: 5
	fmt.Println(countBadPairs([]int{1, 2, 3, 4, 5})) // Output: 0
}