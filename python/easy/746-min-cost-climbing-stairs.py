"""
746. Min Cost Climbing Stairs

You are given an integer array cost where cost[i] is the cost of ith step on a staircase. Once 
you pay the cost, you can either climb one or two steps.

You can either start from the step with index 0, or the step with index 1.

Return the minimum cost to reach the top of the floor.

Example 1:
    Input: cost = [10,15,20]
    Output: 15
    Explanation: You will start at index 1.
        - Pay 15 and climb two steps to reach the top.
        The total cost is 15.

Example 2:
    Input: cost = [1,100,1,1,1,100,1,1,100,1]
    Output: 6
    Explanation: You will start at index 0.
        - Pay 1 and climb two steps to reach index 2.
        - Pay 1 and climb two steps to reach index 4.
        - Pay 1 and climb two steps to reach index 6.
        - Pay 1 and climb one step to reach index 7.
        - Pay 1 and climb two steps to reach index 9.
        - Pay 1 and climb one step to reach the top.
        The total cost is 6.
        

Constraints:
    2 <= cost.length <= 1000
    0 <= cost[i] <= 999
 
Seen this question in a real interview before?
"""

#Naive Recursive
class Solution1:
    def minCostClimbingStairs(self, cost):
        def dfs(i):
            if i < 0: 
                return 0
            if i == 0 or i == 1: 
                return 0
            
            return min(dfs(i-1) + cost[i-1], dfs(i-2) + cost[i-2])
        
        return dfs(len(cost))



#Memoization (Top-Down)
class Solution2:
    def minCostClimbingStairs(self, cost):
        n = len(cost)
        memo = {}

        def dfs(i):
            if i <= 1:
                return 0
            if i in memo:
                return memo[i]
            
            memo[i] = min(dfs(i-1) + cost[i-1], dfs(i-2) + cost[i-2])

            return memo[i]

        return dfs(n)



#Tabulation (Bottom-Up)
class Solution3:
    def minCostClimbingStairs(self, cost):
        n = len(cost)
        dp = [0] * (n+1)  
         
        for i in range(2, n+1):
            dp[i] = min(dp[i-1] + cost[i-1], dp[i-2] + cost[i-2])

        return dp[n]

#Space Optimized (Rolling Variables)
class Solution4:
    def minCostClimbingStairs(self, cost):
        n = len(cost)
        prev2, prev1 = 0, 0

        for i in range(2, n+1):
            curr = min(prev1 + cost[i-1], prev2 + cost[i-2])
            prev2, prev1 = prev1, curr

        return prev1