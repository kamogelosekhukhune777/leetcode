package main

import "fmt"

//https://leetcode.com/problems/generate-parentheses/description
/*
Given n pairs of parentheses, write a function to generate all combinations
of well-formed parentheses.

Example 1:

	Input: n = 3
	Output: ["((()))","(()())","(())()","()(())","()()()"]

Example 2:

	Input: n = 1
	Output: ["()"]


Constraints:

	1 <= n <= 8
*/

// Function to generate all combinations of well-formed parentheses
func generateParenthesis(n int) []string {
	var result []string // Initialize the result list

	// Backtracking function to generate combinations
	var backtrack func(current string, open, close int)
	backtrack = func(current string, open, close int) {
		// Base case: if the current combination length is 2 * n, add to result
		if len(current) == 2*n {
			result = append(result, current)
			return
		}

		// Add an opening parenthesis if the count of open is less than n
		if open < n {
			backtrack(current+"(", open+1, close)
		}

		// Add a closing parenthesis if the count of close is less than open
		if close < open {
			backtrack(current+")", open, close+1)
		}
	}

	// Start the backtracking process with an empty combination, 0 open, and 0 close
	backtrack("", 0, 0)
	return result
}

// Main function to test the generateParenthesis function
func main() {
	n1 := 3
	fmt.Println(generateParenthesis(n1)) // Output: ["((()))","(()())","(())()","()(())","()()()"]

	n2 := 1
	fmt.Println(generateParenthesis(n2)) // Output: ["()"]
}
