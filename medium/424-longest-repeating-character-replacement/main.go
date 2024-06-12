package main

import "fmt"

/*
You are given a string s and an integer k. You can choose any character of the string
and change it to any other uppercase English character. You can perform this operation at most k times.

Return the length of the longest substring containing the same letter you can get after
performing the above operations.


Example 1:

	Input: s = "ABAB", k = 2
	Output: 4
	Explanation: Replace the two 'A's with two 'B's or vice versa.

Example 2:

	Input: s = "AABABBA", k = 1
	Output: 4
	Explanation: Replace the one 'A' in the middle with 'B' and form "AABBBBA".
	The substring "BBBB" has the longest repeating letters, which is 4.
	There may exists other ways to achieve this answer too.


Constraints:

	1 <= s.length <= 10^5
	s consists of only uppercase English letters.
	0 <= k <= s.length
*/

func characterReplacement(s string, k int) int {
	count := make([]int, 26)

	index := func(c byte) int {
		return int(c - 'A')
	}

	maxCount, left, result := 0, 0, 0

	//iterate over the string with the right end of the window
	for right := 0; right < len(s); right++ {
		//increment the count of the current character
		count[index(s[right])]++

		//Update the MaxCount of the most frequent character in the current window
		if count[index(s[right])] > maxCount {
			maxCount = count[index(s[right])]
		}

		//current window size is right - left + 1
		//if the current window size minus the most frequent character count is greather than k
		//it means we need more than k changes to make all characters the same in the window
		if right-left+1-maxCount > k {
			//decrement the count of the leftmost character and move the left end of the window
			count[index(s[left])]--
			left++
		}

		//Upadte the result with the maximum size of the valid window
		result = max(result, right-left+1)
	}

	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	//Example 1
	s := "ABAB"
	k := 2
	fmt.Println(characterReplacement(s, k)) //Output: 4

	//Example 2
	s = "AABABBA"
	k = 1
	fmt.Println(characterReplacement(s, k)) //Output: 4
}
