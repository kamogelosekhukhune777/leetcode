package main

import "fmt"

/*
A phrase is a palindrome if, after converting all uppercase letters into lowercase letters
and removing all non-alphanumeric characters, it reads the same forward and backward. Alphanumeric
characters include letters and numbers.

Given a string s, return true if it is a palindrome, or false otherwise.

Example 1:
	Input: s = "A man, a plan, a canal: Panama"
	Output: true
	Explanation: "amanaplanacanalpanama" is a palindrome.

Example 2:
	Input: s = "race a car"
	Output: false
	Explanation: "raceacar" is not a palindrome.

Example 3:
	Input: s = " "
	Output: true
	Explanation: s is an empty string "" after removing non-alphanumeric characters.
	Since an empty string reads the same forward and backward, it is a palindrome.

Constraints:
	1 <= s.length <= 2 * 105
	s consists only of printable ASCII characters.
*/

// isAlphanumeric checks if a byte is alphanumeric (a-z, A-Z, 0-9)
func isAlphanumeric(char byte) bool {
	return (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') || (char >= '0' && char <= '9')
}

// toLower converts an uppercase letter to lowercase
func toLower(char byte) byte {
	if char >= 'A' && char <= 'Z' {
		return char + 'a' - 'A'
	}
	return char
}

// isPalindrome checks if a string is a palindrome after converting to lowercase and removing non-alphanumeric characters
func isPalindrome(s string) bool {
	// Convert the string to lowercase and remove non-alphanumeric characters
	var modifiedString []byte
	for i := 0; i < len(s); i++ {
		if isAlphanumeric(s[i]) {
			modifiedString = append(modifiedString, toLower(s[i]))
		}
	}

	// Initialize fast and slow pointers
	fast, slow := 0, 0

	// Move fast pointer to the end and slow pointer to the middle
	for fast < len(modifiedString)/2 {
		fast++
		slow++
	}

	// Adjust slow pointer for odd-length strings
	if len(modifiedString)%2 == 1 {
		slow++
	}

	// Compare characters using fast and slow pointers
	for slow < len(modifiedString) {
		if modifiedString[fast] != modifiedString[slow] {
			return false
		}
		fast--
		slow++
	}

	return true
}

func main() {
	// Test cases
	fmt.Println(isPalindrome("A man, a plan, a canal: Panama")) // Output: true
	fmt.Println(isPalindrome("race a car"))                     // Output: false
	fmt.Println(isPalindrome(" "))                              // Output: true
}

/*
func isAlphanumeric(char byte) bool {
	return (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') || (char >= '0' && char <= '9')
}

func toLower(char byte) byte {
	if char >= 'A' && char <= 'Z' {
		return char + 'a' - 'A'
	}
	return char
}

func isPalindrome(s string) bool {
	// Convert the string to lowercase and remove non-alphanumeric characters
	var modifiedString []byte
	for i := 0; i < len(s); i++ {
		if isAlphanumeric(s[i]) {
			modifiedString = append(modifiedString, toLower(s[i]))
		}
	}

	// Initialize pointers
	left, right := 0, len(modifiedString)-1

	// Move pointers towards each other until they meet or cross
	for left < right {
		if modifiedString[left] != modifiedString[right] {
			return false
		}
		left++
		right--
	}

	return true
}
*/
