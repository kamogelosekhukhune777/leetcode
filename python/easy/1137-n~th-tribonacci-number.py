"""
1137. N-th Tribonacci Number

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
    

"""
#Memoization
#Time: O(n), Space: O(n)
class Solution1(object):
    def tribonacci(self, n):
        memo = {}
        
        def dp(k):
            if k in memo:
                return memo[k]
            if k == 0:
                return 0
            if k == 1 or k == 2:
                return 1
            
            memo[k] = dp(k-1) + dp(k-2) + dp(k-3)
            return memo[k]
        
        return dp(n)

#Tabulation:: Time: O(n), Space: O(n).
class Solution2(object):
    def tribonacci(n):
        if n == 0: 
            return 0
        if n == 1 or n == 2: 
            return 1
        
        dp = [0]*(n+1)
        dp[0], dp[1], dp[2] = 0, 1, 1
        
        for i in range(3, n+1):
            dp[i] = dp[i-1] + dp[i-2] + dp[i-3]
        
        return dp[n]

#Tabulation:: Time: O(n), Space: O(1)
class Solution3(object):
    def tribonacci(self, n):
        if n == 0:
            return 0
        if n == 1 or n == 2: 
            return 1
        
        a, b, c = 0, 1, 1
        for i in range(3, n+1):
            a, b, c = b, c, a+b+c
        return c
        