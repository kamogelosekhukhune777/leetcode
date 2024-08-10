package main

import (
	"fmt"
)

/*
In an alien language, surprisingly, they also use English lowercase letters, but possibly in a different order. The
order of the alphabet is some permutation of lowercase letters.

Given a sequence of words written in the alien language, and the order of the alphabet, return true if and only if the
given words are sorted lexicographically in this alien language.

Example 1:

	Input: words = ["hello","leetcode"], order = "hlabcdefgijkmnopqrstuvwxyz"
	Output: true
	Explanation: As 'h' comes before 'l' in this language, then the sequence is sorted.

Example 2:

	Input: words = ["word","world","row"], order = "worldabcefghijkmnpqstuvxyz"
	Output: false
	Explanation: As 'd' comes after 'l' in this language, then words[0] > words[1], hence
		the sequence is unsorted.

Example 3:

	Input: words = ["apple","app"], order = "abcdefghijklmnopqrstuvwxyz"
	Output: false
	Explanation: The first three characters "app" match, and the second string is shorter (in size.) According to
		lexicographical rules "apple" > "app", because 'l' > '∅', where '∅' is defined as the blank character
		which is less than any other character (More info).

Constraints:

	1 <= words.length <= 100
	1 <= words[i].length <= 20
	order.length == 26
	All characters in words[i] and order are English lowercase letters.
*/

func isAlienSorted(words []string, order string) bool {
	// Step 1: Create a mapping from character to its index in the order
	orderMap := make(map[byte]int)
	for i := 0; i < len(order); i++ {
		orderMap[order[i]] = i
	}

	// Step 2: Compare each pair of adjacent words
	for i := 0; i < len(words)-1; i++ {
		if !inCorrectOrder(words[i], words[i+1], orderMap) {
			return false
		}
	}
	return true
}

func inCorrectOrder(word1, word2 string, orderMap map[byte]int) bool {
	minLength := min(len(word1), len(word2))
	for i := 0; i < minLength; i++ {
		if word1[i] != word2[i] {
			return orderMap[word1[i]] < orderMap[word2[i]]
		}
	}
	return len(word1) <= len(word2)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	words1 := []string{"hello", "leetcode"}
	order1 := "hlabcdefgijkmnopqrstuvwxyz"
	fmt.Println(isAlienSorted(words1, order1)) // Output: true

	words2 := []string{"word", "world", "row"}
	order2 := "worldabcefghijkmnpqstuvxyz"
	fmt.Println(isAlienSorted(words2, order2)) // Output: false

	words3 := []string{"apple", "app"}
	order3 := "abcdefghijklmnopqrstuvwxyz"
	fmt.Println(isAlienSorted(words3, order3)) // Output: false
}
