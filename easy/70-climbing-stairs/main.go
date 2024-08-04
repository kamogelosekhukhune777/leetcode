package main

import "fmt"

/*
You are climbing a staircase. It takes n steps to reach the top.

Each time you can either climb 1 or 2 steps. In how many distinct ways can you climb to the top?


Example 1:

Input: n = 2
Output: 2
Explanation: There are two ways to climb to the top.
	1. 1 step + 1 step
	2. 2 steps

Example 2:

Input: n = 3
Output: 3
Explanation: There are three ways to climb to the top.
	1. 1 step + 1 step + 1 step
	2. 1 step + 2 steps
	3. 2 steps + 1 step


Constraints:

	1 <= n <= 45
*/

// climbStairs calculates the number of distinct ways to climb to the top
func climbStairs(n int) int {
	// If n is less than or equal to 2, return n
	if n <= 2 {
		return n
	}

	// Initialize the dp array with n+1 elements
	dp := make([]int, n+1)

	// Base cases
	dp[0] = 1
	dp[1] = 1
	dp[2] = 2

	// Fill the dp array
	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	// The result is the number of ways to reach the nth step
	return dp[n]
}

func main() {
	// Example 1
	n1 := 2
	fmt.Println(climbStairs(n1)) // Output: 2

	// Example 2
	n2 := 3
	fmt.Println(climbStairs(n2)) // Output: 3

	// Additional test cases
	n3 := 5
	fmt.Println(climbStairs(n3)) // Output: 8 (11111, 1112, 1121, 1211, 2111, 221, 212, 122)
}
