"""
198. House Robber

You are a professional robber planning to rob houses along a street. Each house has a certain 
amount of money stashed, the only constraint stopping you from robbing each of them is that 
adjacent houses have security systems connected and it will automatically contact the police 
if two adjacent houses were broken into on the same night.

Given an integer array nums representing the amount of money of each house, return the maximum 
amount of money you can rob tonight without alerting the police.


Example 1:
    Input: nums = [1,2,3,1]
    Output: 4
    Explanation: Rob house 1 (money = 1) and then rob house 3 (money = 3).
        Total amount you can rob = 1 + 3 = 4.

Example 2:
    Input: nums = [2,7,9,3,1]
    Output: 12
    Explanation: Rob house 1 (money = 2), rob house 3 (money = 9) and rob house 5 (money = 1).
        Total amount you can rob = 2 + 9 + 1 = 12.
 

Constraints:
    1 <= nums.length <= 100
    0 <= nums[i] <= 400
"""

#Naive Recursion (Exponential)
#Correct, but exponential O(2^n).
class Solution1(object):
    def rob(self, nums):
        def dfs(i):
            if i < 0: 
                return 0
            
            return max(nums[i] + dfs(i-2), dfs(i-1))
        
        return dfs(len(nums)-1)



#Recursion + Memoization (Top-Down)
#O(n) time, O(n) space for memo.
class Solution2(object):
    def rob(self, nums):
        memo = {}
        def dfs(i):
            if i < 0: 
                return 0
            if i in memo: 
                return memo[i]
            
            memo[i] = max(nums[i] + dfs(i-2), dfs(i-1))

            return memo[i]
        
        return dfs(len(nums)-1)



#Tabulation (Bottom-Up DP)
#O(n) time, O(n) space.
class Solution3(object):
    def rob(self, nums):
        n = len(nums)
        if n == 1: 
            return nums[0]
        
        dp = [0]*n
        dp[0] = nums[0]
        dp[1] = max(nums[0], nums[1])

        for i in range(2, n):
            dp[i] = max(nums[i] + dp[i-2], dp[i-1])

        return dp[-1]



#Space-Optimized DP (Two Variables)
#O(n) time, O(1) space.
class Solution4(object):
    def rob(self, nums):
        prev2, prev1 = 0, 0

        for num in nums:
            curr = max(num + prev2, prev1)
            prev2, prev1 = prev1, curr

        return prev1

