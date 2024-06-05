package main

import "fmt"

/*
Given a string s, return true if the s can be palindrome after deleting at
most one character from it.


Example 1:
	Input: s = "aba"
	Output: true

Example 2:
	Input: s = "abca"
	Output: true
	Explanation: You could delete the character 'c'.

Example 3:
	Input: s = "abc"
	Output: false


Constraints:
	1 <= s.length <= 105
	s consists of lowercase English letters.
*/
// Helper function to check if a substring is a palindrome
func isPalindrome(s string, left int, right int) bool {
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}

// Function to check if the string can be a palindrome after deleting at most one character
func validPalindrome(s string) bool {
	left, right := 0, len(s)-1

	for left < right {
		if s[left] != s[right] {
			// Check the two possibilities: removing the character at left or at right
			return isPalindrome(s, left+1, right) || isPalindrome(s, left, right-1)
		}
		left++
		right--
	}

	return true
}

func main() {
	fmt.Println(validPalindrome("aba"))
	fmt.Println(validPalindrome("abca"))
	fmt.Println(validPalindrome("abc"))
}
