package main

import (
	"fmt"
)

/*
You are given an integer array coins representing coins of different denominations and an integer
amount representing a total amount of money.

Return the fewest number of coins that you need to make up that amount. If that amount of money cannot
be made up by any combination of the coins, return -1.

You may assume that you have an infinite number of each kind of coin.


Example 1:

	Input: coins = [1,2,5], amount = 11
	Output: 3
	Explanation: 11 = 5 + 5 + 1

Example 2:

	Input: coins = [2], amount = 3
	Output: -1

Example 3:

	Input: coins = [1], amount = 0
	Output: 0


Constraints:

	1 <= coins.length <= 12
	1 <= coins[i] <= 2^31 - 1
	0 <= amount <= 10^4
*/

// Function to find the minimum number of coins needed to make up the amount
func coinChange(coins []int, amount int) int {
	// Initialize the DP array with a value greater than any possible number of coins
	dp := make([]int, amount+1)
	for i := range dp {
		dp[i] = amount + 1
	}
	dp[0] = 0

	// Fill the DP array
	for _, coin := range coins {
		for i := coin; i <= amount; i++ {
			dp[i] = min(dp[i], dp[i-coin]+1)
		}
	}

	// Return the result
	if dp[amount] > amount {
		return -1
	}
	return dp[amount]
}

// Helper function to get the minimum of two integers
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	// Example 1
	coins1 := []int{1, 2, 5}
	amount1 := 11
	fmt.Println(coinChange(coins1, amount1)) // Output: 3

	// Example 2
	coins2 := []int{2}
	amount2 := 3
	fmt.Println(coinChange(coins2, amount2)) // Output: -1

	// Example 3
	coins3 := []int{1}
	amount3 := 0
	fmt.Println(coinChange(coins3, amount3)) // Output: 0
}
