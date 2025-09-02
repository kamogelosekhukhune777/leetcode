"""
509. Fibonacci Number

The Fibonacci numbers, commonly denoted F(n) form a sequence, called the Fibonacci sequence, such 
that each number is the sum of the two preceding ones, starting from 0 and 1. That is,

    F(0) = 0, F(1) = 1
    F(n) = F(n - 1) + F(n - 2), for n > 1.
Given n, calculate F(n).

Example 1:
    Input: n = 2
    Output: 1
    Explanation: F(2) = F(1) + F(0) = 1 + 0 = 1.

Example 2:
    Input: n = 3
    Output: 2
    Explanation: F(3) = F(2) + F(1) = 1 + 1 = 2.

Example 3:
    Input: n = 4
    Output: 3
    Explanation: F(4) = F(3) + F(2) = 2 + 1 = 3.
    

Constraints:
    0 <= n <= 30

"""

#Top-Down DP (Memoization)
#Time: O(n), Space: O(n) (recursion + dictionary).
class Solution(object):
    def __init__(self):
        self.memo = {}
        
    def fib(self, n):
        if n in self.memo:
            return self.memo[n]
        
        if n <= 1:
            return n
        
        self.memo[n] = self.fib(n-1) + self.fib(n-2)
        return self.memo[n]



#Bottom-Up DP (Tabulation)
#Time: O(n), Space: O(n).
class Solution2(object):
    def fib(self, n):
        if n <= 1:
            return n
        
        dp = [0] * (n+1)
        dp[0], dp[1] = 0, 1
        
        for i in range(2, n+1):
            dp[i] = dp[i-1] + dp[i-2]
        
        return dp[n]



#Space Optimized (Two Variables)
# Time: O(n), Space: O(1).
class Solution3(object):
    def fib(self, n):
        if n <= 1:
            return n
        
        prev2, prev1 = 0, 1
        for _ in range(2, n+1):
            curr = prev1 + prev2
            prev2, prev1 = prev1, curr
        
        return prev1
