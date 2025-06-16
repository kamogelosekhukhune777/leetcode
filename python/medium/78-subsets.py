"""
Given an integer array nums of unique elements, return all 
possible subsets (the power set).

The solution set must not contain duplicate subsets. Return the 
solution in any order.

Example 1:
    Input: nums = [1,2,3]
    Output: [[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]

Example 2:
    Input: nums = [0]
    Output: [[],[0]]
 
Constraints:
    1 <= nums.length <= 10
    -10 <= nums[i] <= 10
    All the numbers of nums are unique.
"""

#Iterative Expansion
class Solution1(object):
    def subsets(self, nums):
        result = [[]]
        
        for num in nums:
            new_subsets = []
            for subset in result:
                new_subsets.append(subset + [num])
            result.extend(new_subsets)
        
        return result
    
#Backtracking(DFS)
class Solution2(object):
    def subsets(self,nums):
        result = []
        
        def backtrack(start, path):
            result.append(path[:])
            for i in range(start, len(nums)):
                path.append(nums[i])
                backtrack(i + 1, path)
                path.pop()
        
        backtrack(0, [])
        return result
    

#Bit-Manipulation
class Solution3(object):
    def subsets(self,nums):
        n = len(nums)
        result = []

        for bitmask in range(1 << n):  # 0 to 2^n - 1
            subset = []
            for i in range(n):
                if bitmask & (1 << i):
                    subset.append(nums[i])
            result.append(subset)

        return result
