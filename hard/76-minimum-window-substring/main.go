package main

import (
	"fmt"
	"math"
)

/*
Given two strings s and t of lengths m and n respectively, return the minimum window

substring of s such that every character in t (including duplicates) is included in
the window. If there is no such substring, return the empty string "".
The testcases will be generated such that the answer is unique.

Example 1:

	Input: s = "ADOBECODEBANC", t = "ABC"
	Output: "BANC"
	Explanation: The minimum window substring "BANC" includes 'A', 'B', and 'C' from string t.

Example 2:

	Input: s = "a", t = "a"
	Output: "a"
	Explanation: The entire string s is the minimum window.

Example 3:

	Input: s = "a", t = "aa"
	Output: ""
	Explanation: Both 'a's from t must be included in the window.
	Since the largest window of s only has one 'a', return empty string.

Constraints:

	m == s.length
	n == t.length
	1 <= m, n <= 105
	s and t consist of uppercase and lowercase English letters.


Follow up: Could you find an algorithm that runs in O(m + n) time?
*/

func minWindow(s string, t string) string {
	//if t is longer than s, its impossible to find a valid window
	if len(t) > len(s) {
		return ""
	}

	//to store the countt of each character in t
	charCountT := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		charCountT[t[i]]++
	}

	//to keep track of the window
	left, right := 0, 0
	required := len(charCountT)
	formed := 0
	windowCounts := make(map[byte]int)

	//to keep track of the smallest window
	minLen := math.MaxInt32
	start, end := 0, 0

	//move right pointer to expand the window
	for right < len(s) {
		char := s[right]
		windowCounts[char]++

		if _, ok := charCountT[char]; ok && windowCounts[char] == charCountT[char] {
			formed++
		}
	}

	//try to contact the window ubtil it ceases to be desirable
	for left <= right && formed == required {
		char := s[left]

		//save the smallest window
		if right-left+1 < minLen {
			minLen = right - left + 1
			start = left
			end = right
		}

		//the character at the position pointed by the left pointer is no longer a part of the window
		if _, ok := charCountT[char]; ok && windowCounts[char] < charCountT[char] {
			formed--

			//move the left pointer ahead to look for a new window
			left++
		}

		//keep expanding the window by moving the right pointer
		right++
	}

	if minLen == math.MaxInt32 {
		return ""
	}

	return s[start : end+1]
}

func main() {

	//Example 1
	s := "ADOBECODEBANC"
	t := "ABC"
	fmt.Println(minWindow(s, t)) //Output: "BANC"

	//Example 2

	s = "a"
	t = "a"
	fmt.Println(minWindow(s, t)) //Output: "a"

	//Example 3
	s = "a"
	t = "aa"
	fmt.Println(minWindow(s, t)) //Output: ""
}
