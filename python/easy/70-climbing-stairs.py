"""
70. Climbing Stairs

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
"""

# Naïve Recursion
# Time = O(2^n), too slow.
class Solution1(object):
    def climbStairs(self, n):
        if n <= 2:
            return n
        return self.climbStairs(n-1) + self.climbStairs(n-2)



#Memoization (Top-Down)
#Time = O(n), Space = O(n) (stack + dictionary).
class Solution2(object):
    def __init__(self):
        self.memo = {}
        
    def climbStairs(self, n):
        if n in self.memo:
            return self.memo[n]
        
        if n <= 2:
            return n
        
        self.memo[n] = self.climbStairs(n-1) + self.climbStairs(n-2)
        return self.memo[n]



#Tabulation (Bottom-Up)
#Time = O(n), Space = O(n).
class Solution3(object):
    def climbStairs(self, n):
        if n <= 2:
            return n
        
        dp = [0] * (n+1)
        dp[1], dp[2] = 1, 2
        
        for i in range(3, n+1):
            dp[i] = dp[i-1] + dp[i-2]
        
        return dp[n]



#Space Optimized (Two Variables)
#Time = O(n), Space = O(1).
class Solution(object):
    def climbStairs(self, n):
        if n <= 2:
            return n
        
        prev2, prev1 = 1, 2
        for _ in range(3, n+1):
            curr = prev1 + prev2
            prev2, prev1 = prev1, curr
        
        return prev1

