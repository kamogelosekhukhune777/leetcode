package main

import (
	"fmt"
)

/*
You have intercepted a secret message encoded as a string of numbers. The message is decoded
via the following mapping:

	"1" -> 'A'

	"2" -> 'B'

	...

	"25" -> 'Y'

	"26" -> 'Z'

However, while decoding the message, you realize that there are many different ways you can decode the message
because some codes are contained in other codes ("2" and "5" vs "25").

For example, "11106" can be decoded into:

	"AAJF" with the grouping (1, 1, 10, 6)
	"KJF" with the grouping (11, 10, 6)
	The grouping (1, 11, 06) is invalid because "06" is not a valid code (only "6" is valid).

Note: there may be strings that are impossible to decode.

Given a string s containing only digits, return the number of ways to decode it. If the entire string cannot be
decoded in any valid way, return 0.

The test cases are generated so that the answer fits in a 32-bit integer.

Example 1:

	Input: s = "12"
	Output: 2
	Explanation:
		"12" could be decoded as "AB" (1 2) or "L" (12).

Example 2:

	Input: s = "226"
	Output: 3
	Explanation:
		"226" could be decoded as "BZ" (2 26), "VF" (22 6), or "BBF" (2 2 6).

Example 3:

	Input: s = "06"
	Output: 0
	Explanation:
		"06" cannot be mapped to "F" because of the leading zero ("6" is different from "06"). In this case, the string
		is not a valid encoding, so return 0.

Constraints:

	1 <= s.length <= 100
	s contains only digits and may contain leading zero(s).
*/
func numDecodings(s string) int {
	// If the input string is empty, there are no ways to decode it
	if len(s) == 0 {
		return 0
	}

	// Create a dp array to store the number of ways to decode up to each index
	dp := make([]int, len(s)+1)

	// Base case: there is one way to decode an empty string
	dp[0] = 1

	// If the first character is not '0', there is one way to decode it
	if s[0] != '0' {
		dp[1] = 1
	} else {
		dp[1] = 0
	}

	// Fill the dp array
	for i := 2; i <= len(s); i++ {
		// Check if the single character substring is valid
		if s[i-1] != '0' {
			dp[i] += dp[i-1]
		}
		// Check if the two-character substring is valid
		twoDigit := (s[i-2]-'0')*10 + (s[i-1] - '0')
		if twoDigit >= 10 && twoDigit <= 26 {
			dp[i] += dp[i-2]
		}
	}

	// The result is the number of ways to decode the entire string
	return dp[len(s)]
}

func main() {
	// Example 1
	s1 := "12"
	fmt.Println(numDecodings(s1)) // Output: 2

	// Example 2
	s2 := "226"
	fmt.Println(numDecodings(s2)) // Output: 3

	// Example 3
	s3 := "06"
	fmt.Println(numDecodings(s3)) // Output: 0
}
