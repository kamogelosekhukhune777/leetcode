"""
416. Partition Equal Subset Sum

Given an integer array nums, return true if you can partition the array into two subsets such 
that the sum of the elements in both subsets is equal or false otherwise.


Example 1:
    Input: nums = [1,5,11,5]
    Output: true
    Explanation: The array can be partitioned as [1, 5, 5] and [11].

Example 2:
    Input: nums = [1,2,3,5]
    Output: false
    Explanation: The array cannot be partitioned into equal sum subsets.
 

Constraints:
    1 <= nums.length <= 200
    1 <= nums[i] <= 100

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