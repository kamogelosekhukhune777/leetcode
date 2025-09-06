"""
494. Target Sum

You are given an integer array nums and an integer target.

You want to build an expression out of nums by adding one of the symbols '+' and '-' before each 
integer in nums and then concatenate all the integers.

For example, if nums = [2, 1], you can add a '+' before 2 and a '-' before 1 and concatenate them 
to build the expression "+2-1".
Return the number of different expressions that you can build, which evaluates to target.

Example 1:
    Input: nums = [1,1,1,1,1], target = 3
    Output: 5
    Explanation: There are 5 ways to assign symbols to make the sum of nums be target 3.
        -1 + 1 + 1 + 1 + 1 = 3
        +1 - 1 + 1 + 1 + 1 = 3
        +1 + 1 - 1 + 1 + 1 = 3
        +1 + 1 + 1 - 1 + 1 = 3
        +1 + 1 + 1 + 1 - 1 = 3

Example 2:
    Input: nums = [1], target = 1
    Output: 1
 

Constraints:
    1 <= nums.length <= 20
    0 <= nums[i] <= 1000
    0 <= sum(nums[i]) <= 1000
    -1000 <= target <= 1000
"""

#Naive Recursion (exponential)
class Solution1:
    def findTargetSumWays(self, nums, target):
        n = len(nums)

        def dfs(i, total):
            if i == n:
                return 1 if total == target else 0
            return dfs(i+1, total + nums[i]) + dfs(i+1, total - nums[i])

        return dfs(0, 0)

#Recursion + Memoization
class Solution2:
    def findTargetSumWays(self, nums, target):
        n = len(nums)
        memo = {}

        def dfs(i, total):
            if i == n:
                return 1 if total == target else 0
            if (i, total) in memo:
                return memo[(i, total)]

            memo[(i, total)] = dfs(i+1, total + nums[i]) + dfs(i+1, total - nums[i])
            return memo[(i, total)]

        return dfs(0, 0)

#Subset-Sum Transformation (Tabulation, 1D DP)
class Solution3:
    def findTargetSumWays(self, nums, target):
        total = sum(nums)
        if (target + total) % 2 != 0 or abs(target) > total:
            return 0

        subset_sum = (target + total) // 2
        dp = [0] * (subset_sum + 1)
        dp[0] = 1

        for num in nums:
            for s in range(subset_sum, num-1, -1):  # reverse to prevent reuse
                dp[s] += dp[s - num]

        return dp[subset_sum]