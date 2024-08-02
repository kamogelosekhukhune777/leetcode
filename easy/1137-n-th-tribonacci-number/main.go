package main

import "fmt"

/*
The Tribonacci sequence Tn is defined as follows:

T0 = 0, T1 = 1, T2 = 1, and Tn+3 = Tn + Tn+1 + Tn+2 for n >= 0.

Given n, return the value of Tn.


Example 1:

	Input: n = 4
	Output: 4
	Explanation:
		T_3 = 0 + 1 + 1 = 2
		T_4 = 1 + 1 + 2 = 4

Example 2:

	Input: n = 25
	Output: 1389537


Constraints:

	0 <= n <= 37
	The answer is guaranteed to fit within a 32-bit integer, ie. answer <= 2^31 - 1.
*/

// Function to calculate the n-th Tribonacci number
func tribonacci(n int) int {
	// Handle base cases
	if n == 0 {
		return 0
	}
	if n == 1 || n == 2 {
		return 1
	}

	// Initialize the DP array
	dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = 1
	dp[2] = 1

	// Fill the DP array using the recurrence relation
	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2] + dp[i-3]
	}

	// Return the n-th Tribonacci number
	return dp[n]
}

func main() {
	// Example 1
	n1 := 4
	fmt.Println(tribonacci(n1)) // Output: 4

	// Example 2
	n2 := 25
	fmt.Println(tribonacci(n2)) // Output: 1389537
}
