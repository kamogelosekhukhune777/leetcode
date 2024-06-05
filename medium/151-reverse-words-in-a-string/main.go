package main

import (
	"fmt"
	"strings"
)

/*
Given an input string s, reverse the order of the words.

A word is defined as a sequence of non-space characters. The words in s will be separated by at least one space.

Return a string of the words in reverse order concatenated by a single space.

Note that s may contain leading or trailing spaces or multiple spaces between two words. The returned string should
only have a single space separating the words. Do not include any extra spaces.


Example 1:

	Input: s = "the sky is blue"
	Output: "blue is sky the"

Example 2:

	Input: s = "  hello world  "
	Output: "world hello"
	Explanation: Your reversed string should not contain leading or trailing spaces.

Example 3:

	Input: s = "a good   example"
	Output: "example good a"
	Explanation: You need to reduce multiple spaces between two words to a single space in the reversed string.

Constraints:

	1 <= s.length <= 104
	s contains English letters (upper-case and lower-case), digits, and spaces ' '.
	 There is at least one word in s.


Follow-up: If the string data type is mutable in your language, can you solve it in-place with O(1) extra space?
*/

// Function to reverse the order of words in the input string
func reverseWords(s string) string {
	// Trim leading and trailing spaces
	s = strings.TrimSpace(s)

	// Split the string by spaces into a slice of words
	words := strings.Fields(s)

	// Reverse the order of the words
	left, right := 0, len(words)-1
	for left < right {
		words[left], words[right] = words[right], words[left]
		left++
		right--
	}

	// Join the reversed words with a single space
	return strings.Join(words, " ")
}

func main() {
	s1 := "the sky is blue"
	fmt.Println(reverseWords(s1)) // Output: "blue is sky the"

	s2 := "  hello world  "
	fmt.Println(reverseWords(s2)) // Output: "world hello"

	s3 := "a good   example"
	fmt.Println(reverseWords(s3)) // Output: "example good a"
}
