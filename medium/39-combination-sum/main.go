package main

import (
	"fmt"
	"sort"
)

/*
Given an array of distinct integers candidates and a target integer target, return a list of all unique combinations
of candidates where the chosen numbers sum to target. You may return the combinations in any order.

The same number may be chosen from candidates an unlimited number of times. Two combinations are unique if the

frequency of at least one of the chosen numbers is different.

The test cases are generated such that the number of unique combinations that sum up to target is less than 150
combinations for the given input.

Example 1:

	Input: candidates = [2,3,6,7], target = 7
	Output: [[2,2,3],[7]]
	Explanation:
		2 and 3 are candidates, and 2 + 2 + 3 = 7. Note that 2 can be used multiple times.
		7 is a candidate, and 7 = 7.
		These are the only two combinations.

Example 2:

	Input: candidates = [2,3,5], target = 8
	Output: [[2,2,2,2],[2,3,3],[3,5]]

Example 3:

	Input: candidates = [2], target = 1
	Output: []


Constraints:

	1 <= candidates.length <= 30
	2 <= candidates[i] <= 40
	All elements of candidates are distinct.
	1 <= target <= 40
*/
// Function to find all unique combinations that sum up to the target
func combinationSum(candidates []int, target int) [][]int {
	var result [][]int
	var combination []int

	// Sort the candidates to help with pruning the search space
	sort.Ints(candidates)

	// Helper function for backtracking
	var backtrack func(start, target int)
	backtrack = func(start, target int) {
		if target == 0 {
			// If the target is zero, we found a valid combination
			combinationCopy := make([]int, len(combination))
			copy(combinationCopy, combination)
			result = append(result, combinationCopy)
			return
		}

		for i := start; i < len(candidates); i++ {
			if candidates[i] > target {
				// No need to continue if the candidate is greater than the target
				break
			}
			// Include the candidate
			combination = append(combination, candidates[i])
			// Recurse with the updated target
			backtrack(i, target-candidates[i])
			// Backtrack: remove the last candidate
			combination = combination[:len(combination)-1]
		}
	}

	// Start the backtracking process from the first candidate
	backtrack(0, target)
	return result
}

func main() {
	// Example 1
	candidates1 := []int{2, 3, 6, 7}
	target1 := 7
	fmt.Println(combinationSum(candidates1, target1)) // Output: [[2, 2, 3], [7]]

	// Example 2
	candidates2 := []int{2, 3, 5}
	target2 := 8
	fmt.Println(combinationSum(candidates2, target2)) // Output: [[2, 2, 2, 2], [2, 3, 3], [3, 5]]

	// Example 3
	candidates3 := []int{2}
	target3 := 1
	fmt.Println(combinationSum(candidates3, target3)) // Output: []
}
