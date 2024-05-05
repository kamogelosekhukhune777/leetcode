package main

import "fmt"

/*
We define the usage of capitals in a word to be right when one of the following cases holds:
	All letters in this word are capitals, like "USA".
	All letters in this word are not capitals, like "leetcode".
	Only the first letter in this word is capital, like "Google".
Given a string word, return true if the usage of capitals in it is right.


Example 1:
	Input: word = "USA"
    Output: true

Example 2:
   Input: word = "FlaG"
   Output: false


Constraints:
   1 <= word.length <= 100
   word consists of lowercase and uppercase English letters.
*/

//Soloution
/*
Check if all letters in the word are uppercase.
Check if all letters in the word are lowercase.
Check if only the first letter in the word is uppercase.
Here's the implementation in Go:
*/

func detectCapitalUse(word string) bool {
	if isAllUpper(word) {
		return true
	}

	if isAllLower(word) {
		return true
	}

	if len(word) > 0 && word[0] >= 'A' && word[0] <= 'Z' {
		for i := 1; i < len(word); i++ {
			if word[i] < 'a' || word[i] > 'z' {
				return false
			}
		}
		return true
	}

	return false
}

func isAllUpper(word string) bool {
	for _, char := range word {
		if char < 'A' || char > 'Z' {
			return false
		}
	}

	return true
}

func isAllLower(word string) bool {
	for _, char := range word {
		if char < 'a' || char > 'z' {
			return false
		}
	}

	return true
}

func main() {
	word1 := "USA"
	word2 := "FlaG"

	fmt.Println("is there correct usage of capitalization:", detectCapitalUse(word1))
	fmt.Println("is there correct usage of capitalization:", detectCapitalUse(word2))
}
