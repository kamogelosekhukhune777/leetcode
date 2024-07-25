package main

import (
	"fmt"
)

/*
You are a product manager and currently leading a team to develop a new product. Unfortunately, the latest version of your product fails the quality check. Since each version is developed based on the previous version, all the versions after a bad version are also bad.

Suppose you have n versions [1, 2, ..., n] and you want to find out the first bad one, which causes all the following ones to be bad.

You are given an API bool isBadVersion(version) which returns whether version is bad. Implement a function to find the first bad version. You should minimize the number of calls to the API.



Example 1:

	Input: n = 5, bad = 4
	Output: 4
	Explanation:
		call isBadVersion(3) -> false
		call isBadVersion(5) -> true
		call isBadVersion(4) -> true
		Then 4 is the first bad version.

Example 2:

	Input: n = 1, bad = 1
	Output: 1


Constraints:

	1 <= bad <= n <= 2^31 - 1
*/

// Mock implementation of the isBadVersion API
func isBadVersion(version int) bool {
	// This is a mock implementation, in real scenarios, it would be provided.
	// Assuming bad version starts at 4 for this example
	return version >= 4
}

// Function to find the first bad version using binary search
func firstBadVersion(n int) int {
	low, high := 1, n

	for low < high {
		mid := low + (high-low)/2

		if isBadVersion(mid) {
			high = mid // First bad version is at mid or before mid
		} else {
			low = mid + 1 // First bad version is after mid
		}
	}

	return low // or high, since low == high
}

func main() {
	n1 := 5
	fmt.Printf("First bad version in %d versions: %d\n", n1, firstBadVersion(n1)) // Output: 4

	n2 := 1
	fmt.Printf("First bad version in %d versions: %d\n", n2, firstBadVersion(n2)) // Output: 1
}
