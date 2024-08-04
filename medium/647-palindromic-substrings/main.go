package main

import (
	"fmt"
)

/*
Given a string s, return the number of palindromic substrings in it.

A string is a palindrome when it reads the same backward as forward.

A substring is a contiguous sequence of characters within the string.


Example 1:

	Input: s = "abc"
	Output: 3
	Explanation: Three palindromic strings: "a", "b", "c".

Example 2:

	Input: s = "aaa"
	Output: 6
	Explanation: Six palindromic strings: "a", "a", "a", "aa", "aa", "aaa".


Constraints:

	1 <= s.length <= 1000
	s consists of lowercase English letters.
*/
// Function to count the number of palindromic substrings
func countSubstrings(s string) int {
	n := len(s)
	dp := make([][]bool, n)
	for i := range dp {
		dp[i] = make([]bool, n)
	}
	count := 0

	// Single character substrings
	for i := 0; i < n; i++ {
		dp[i][i] = true
		count++
	}

	// Substrings of length 2
	for i := 0; i < n-1; i++ {
		if s[i] == s[i+1] {
			dp[i][i+1] = true
			count++
		}
	}

	// Substrings of length greater than 2
	for length := 3; length <= n; length++ {
		for i := 0; i <= n-length; i++ {
			j := i + length - 1
			if s[i] == s[j] && dp[i+1][j-1] {
				dp[i][j] = true
				count++
			}
		}
	}

	return count
}

func main() {
	// Example 1
	s1 := "abc"
	fmt.Println(countSubstrings(s1)) // Output: 3

	// Example 2
	s2 := "aaa"
	fmt.Println(countSubstrings(s2)) // Output: 6
}
