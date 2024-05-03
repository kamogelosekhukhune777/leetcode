package main

import "fmt"

/*
Given two string arrays word1 and word2, return true if the two arrays represent the same string, and false otherwise.

A string is represented by an array if the array elements concatenated in order forms the string.


Example 1:

Input: word1 = ["ab", "c"], word2 = ["a", "bc"]
Output: true
Explanation: word1 represents string "ab" + "c" -> "abc"
             word2 represents string "a" + "bc" -> "abc"
             The strings are the same, so return true.

Example 2:

Input: word1 = ["a", "cb"], word2 = ["ab", "c"]
Output: false

Example 3:

Input: word1  = ["abc", "d", "defg"], word2 = ["abcddefg"]
Output: true

Constraints:

   1 <= word1.length, word2.length <= 103
   1 <= word1[i].length, word2[i].length <= 103
   1 <= sum(word1[i].length), sum(word2[i].length) <= 103
   word1[i] and word2[i] consist of lowercase letters.
*/

func arrayStringsAreEqual(word1 []string, word2 []string) bool {
	ptr1, ptr2 := 0, 0
	charPtr1, charPtr2 := 0, 0

	for ptr1 < len(word1) && ptr2 < len(word2) {
		if word1[ptr1][charPtr1] != word2[ptr2][charPtr2] {
			return false
		}

		charPtr1++
		charPtr2++

		if charPtr1 == len(word1[ptr1]) {
			charPtr1 = 0
			ptr1++
		}
		if charPtr2 == len(word2[ptr2]) {
			charPtr2 = 0
			ptr2++
		}
	}

	return ptr1 == len(word1) && ptr2 == len(word2)
}

func main() {
	word1 := []string{"ab", "c"}
	word2 := []string{"a", "bc"}
	fmt.Println(arrayStringsAreEqual(word1, word2))

	word3 := []string{"a", "cb"}
	word4 := []string{"ab", "c"}
	fmt.Println(arrayStringsAreEqual(word3, word4))

	word5 := []string{"abc", "d", "defg"}
	word6 := []string{"abcddefg"}
	fmt.Println(arrayStringsAreEqual(word5, word6))
}
