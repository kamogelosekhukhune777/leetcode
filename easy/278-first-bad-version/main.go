package main

import (
	"fmt"
)

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
