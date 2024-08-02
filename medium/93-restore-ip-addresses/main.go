package main

import (
	"fmt"
	"strconv"
	"strings"
)

/*
A valid IP address consists of exactly four integers separated by single dots. Each integer
is between 0 and 255 (inclusive) and cannot have leading zeros.

For example, "0.1.2.201" and "192.168.1.1" are valid IP addresses, but "0.011.255.245",
"192.168.1.312" and "192.168@1.1" are invalid IP addresses.
Given a string s containing only digits, return all possible valid IP addresses that can be
formed by inserting dots into s. You are not allowed to reorder or remove any digits in s. You
may return the valid IP addresses in any order.


Example 1:

	Input: s = "25525511135"
	Output: ["255.255.11.135","255.255.111.35"]

Example 2:

	Input: s = "0000"
	Output: ["0.0.0.0"]

Example 3:

	Input: s = "101023"
	Output: ["1.0.10.23","1.0.102.3","10.1.0.23","10.10.2.3","101.0.2.3"]


Constraints:

	1 <= s.length <= 20
	s consists of digits only.
*/

func restoreIpAddresses(s string) []string {
	var result []string
	backtrack(s, 0, []string{}, &result)
	return result
}

func backtrack(s string, start int, path []string, result *[]string) {
	// If we have 4 parts and we have consumed all the string, it's a valid IP
	if len(path) == 4 {
		if start == len(s) {
			*result = append(*result, strings.Join(path, "."))
		}
		return
	}

	// Try to place a dot in every possible place and check if the parts are valid
	for i := 1; i <= 3; i++ {
		if start+i > len(s) {
			break
		}
		part := s[start : start+i]
		if isValid(part) {
			backtrack(s, start+i, append(path, part), result)
		}
	}
}

func isValid(part string) bool {
	if len(part) > 1 && part[0] == '0' { // Leading zero is not allowed
		return false
	}
	num, err := strconv.Atoi(part)
	if err != nil || num < 0 || num > 255 { // Part should be between 0 and 255
		return false
	}
	return true
}

func main() {
	// Example 1
	s1 := "25525511135"
	fmt.Println(restoreIpAddresses(s1)) // Output: ["255.255.11.135","255.255.111.35"]

	// Example 2
	s2 := "0000"
	fmt.Println(restoreIpAddresses(s2)) // Output: ["0.0.0.0"]

	// Example 3
	s3 := "101023"
	fmt.Println(restoreIpAddresses(s3)) // Output: ["1.0.10.23","1.0.102.3","10.1.0.23","10.10.2.3","101.0.2.3"]
}
