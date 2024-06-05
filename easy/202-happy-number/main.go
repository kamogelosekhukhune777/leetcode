package main

import "fmt"

/*
Write an algorithm to determine if a number n is happy.

A happy number is a number defined by the following process:

	-Starting with any positive integer, replace the number by the sum of the
	 squares of its digits.
	-Repeat the process until the number equals 1 (where it will stay), or it loops endlessly
	 in a cycle which does not include 1.
	-Those numbers for which this process ends in 1 are happy.
Return true if n is a happy number, and false if not.


Example 1:
	Input: n = 19
	Output: true
	Explanation:
	12 + 92 = 82
	82 + 22 = 68
	62 + 82 = 100
	12 + 02 + 02 = 1
Example 2:
	Input: n = 2
	Output: false


Constraints:

1 <= n <= 231 - 1
*/

// Helper function to calculate the sum of the squares of the digits of a number
func sumOfSquares(n int) int {
	sum := 0
	for n > 0 {
		digit := n % 10
		sum += digit * digit
		n /= 10
	}
	return sum
}

// Function to determine if a number is happy
func isHappy(n int) bool {
	// Initialize slow and fast pointers
	slow := n
	fast := n

	for {
		// Move slow pointer one step
		slow = sumOfSquares(slow)
		// Move fast pointer two steps
		fast = sumOfSquares(sumOfSquares(fast))

		// If fast pointer equals slow pointer, either we've found a cycle or the number is happy
		if slow == fast {
			break
		}
	}

	// Check if we have reached 1
	return slow == 1
}

func main() {
	// Test cases
	fmt.Println(isHappy(19)) // Output: true
	fmt.Println(isHappy(2))  // Output: false
}
