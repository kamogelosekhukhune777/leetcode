package main

import "fmt"

//https://leetcode.com/problems/letter-combinations-of-a-phone-number/description
/*
Given a string containing digits from 2-9 inclusive, return all possible letter
combinations that the number could represent. Return the answer in any order.

A mapping of digits to letters (just like on the telephone buttons) is given
below. Note that 1 does not map to any letters.


Example 1:

	Input: digits = "23"
	Output: ["ad","ae","af","bd","be","bf","cd","ce","cf"]
Example 2:

	Input: digits = ""
	Output: []
Example 3:

	Input: digits = "2"
	Output: ["a","b","c"]


Constraints:

	0 <= digits.length <= 4
	digits[i] is a digit in the range ['2', '9'].
*/
// Function to generate all letter combinations
func letterCombinations(digits string) []string {
	// Mapping of digits to corresponding letters
	digitToLetters := map[byte]string{
		'2': "abc",
		'3': "def",
		'4': "ghi",
		'5': "jkl",
		'6': "mno",
		'7': "pqrs",
		'8': "tuv",
		'9': "wxyz",
	}

	var result []string // Initialize the result list

	// Edge case: if the input is empty, return an empty list
	if len(digits) == 0 {
		return result
	}

	// Backtracking function to generate combinations
	var backtrack func(index int, path string)
	backtrack = func(index int, path string) {
		// Base case: if the current index reaches the length of digits, add the combination to the result
		if index == len(digits) {
			result = append(result, path)
			return
		}

		// Get the letters corresponding to the current digit
		letters := digitToLetters[digits[index]]

		// Iterate through the letters and build combinations
		for i := 0; i < len(letters); i++ {
			// Append the current letter to the path and recursively generate combinations for the next digit
			backtrack(index+1, path+string(letters[i]))
		}
	}

	// Start the backtracking process from the first digit
	backtrack(0, "")
	return result
}

// Main function to test the letterCombinations function
func main() {
	digits1 := "23"
	fmt.Println(letterCombinations(digits1)) // Output: ["ad","ae","af","bd","be","bf","cd","ce","cf"]

	digits2 := ""
	fmt.Println(letterCombinations(digits2)) // Output: []

	digits3 := "2"
	fmt.Println(letterCombinations(digits3)) // Output: ["a","b","c"]
}
