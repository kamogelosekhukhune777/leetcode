package main

import (
	"fmt"
)

/*
Given a string s and a dictionary of strings wordDict, return true if s can be segmented into a space-separated
sequence of one or more dictionary words.

Note that the same word in the dictionary may be reused multiple times in the segmentation.



Example 1:

	Input: s = "leetcode", wordDict = ["leet","code"]
	Output: true
	Explanation: Return true because "leetcode" can be segmented as "leet code".

Example 2:

	Input: s = "applepenapple", wordDict = ["apple","pen"]
	Output: true
	Explanation: Return true because "applepenapple" can be segmented as "apple pen apple".
		Note that you are allowed to reuse a dictionary word.

Example 3:

	Input: s = "catsandog", wordDict = ["cats","dog","sand","and","cat"]
	Output: false


Constraints:

	1 <= s.length <= 300
	1 <= wordDict.length <= 1000
	1 <= wordDict[i].length <= 20
	s and wordDict[i] consist of only lowercase English letters.
	All the strings of wordDict are unique.
*/

// Function to check if the string can be segmented into words from the dictionary
func wordBreak(s string, wordDict []string) bool {
	// Convert wordDict to a set for O(1) look-up time
	wordSet := make(map[string]bool)
	for _, word := range wordDict {
		wordSet[word] = true
	}

	// Initialize the dp array
	dp := make([]bool, len(s)+1)
	dp[0] = true

	// Fill the dp array
	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			if dp[j] && wordSet[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}

	return dp[len(s)]
}

func main() {
	// Example 1
	s1 := "leetcode"
	wordDict1 := []string{"leet", "code"}
	fmt.Println(wordBreak(s1, wordDict1)) // Output: true

	// Example 2
	s2 := "applepenapple"
	wordDict2 := []string{"apple", "pen"}
	fmt.Println(wordBreak(s2, wordDict2)) // Output: true

	// Example 3
	s3 := "catsandog"
	wordDict3 := []string{"cats", "dog", "sand", "and", "cat"}
	fmt.Println(wordBreak(s3, wordDict3)) // Output: false
}
