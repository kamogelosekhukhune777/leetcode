package main

import "fmt"

/*
Given a string s and a dictionary of strings wordDict, add spaces in s to construct a sentence where each word is
a valid dictionary word. Return all such possible sentences in any order.

Note that the same word in the dictionary may be reused multiple times in the segmentation.


Example 1:

	Input: s = "catsanddog", wordDict = ["cat","cats","and","sand","dog"]
	Output: ["cats and dog","cat sand dog"]

Example 2:

	Input: s = "pineapplepenapple", wordDict = ["apple","pen","applepen","pine","pineapple"]
	Output: ["pine apple pen apple","pineapple pen apple","pine applepen apple"]
	Explanation: Note that you are allowed to reuse a dictionary word.

Example 3:

	Input: s = "catsandog", wordDict = ["cats","dog","sand","and","cat"]
	Output: []


Constraints:

	1 <= s.length <= 20
	1 <= wordDict.length <= 1000
	1 <= wordDict[i].length <= 10
	s and wordDict[i] consist of only lowercase English letters.
	All the strings of wordDict are unique.
	Input is generated in a way that the length of the answer doesn't exceed 10^5.

*/

// Function to return all possible sentences formed by adding spaces in s
func wordBreak(s string, wordDict []string) []string {
	wordSet := make(map[string]bool)
	for _, word := range wordDict {
		wordSet[word] = true
	}

	dp := make([][]string, len(s)+1)
	dp[0] = []string{""}

	for i := 1; i <= len(s); i++ {
		var newSentences []string
		for j := 0; j < i; j++ {
			if len(dp[j]) > 0 && wordSet[s[j:i]] {
				for _, sentence := range dp[j] {
					if sentence == "" {
						newSentences = append(newSentences, s[j:i])
					} else {
						newSentences = append(newSentences, sentence+" "+s[j:i])
					}
				}
			}
		}
		dp[i] = newSentences
	}

	return dp[len(s)]
}

func main() {
	// Example 1
	s1 := "catsanddog"
	wordDict1 := []string{"cat", "cats", "and", "sand", "dog"}
	fmt.Println(wordBreak(s1, wordDict1)) // Output: ["cats and dog","cat sand dog"]

	// Example 2
	s2 := "pineapplepenapple"
	wordDict2 := []string{"apple", "pen", "applepen", "pine", "pineapple"}
	fmt.Println(wordBreak(s2, wordDict2)) // Output: ["pine apple pen apple","pineapple pen apple","pine applepen apple"]

	// Example 3
	s3 := "catsandog"
	wordDict3 := []string{"cats", "dog", "sand", "and", "cat"}
	fmt.Println(wordBreak(s3, wordDict3)) // Output: []
}
