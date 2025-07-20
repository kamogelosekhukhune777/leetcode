"""
448. Find All Numbers Disappeared in an Array

Given an array nums of n integers where nums[i] is in the range [1, n], return an 
array of all the integers in the range [1, n] that do not appear in nums.

 
Example 1:
    Input: nums = [4,3,2,7,8,2,3,1]
    Output: [5,6]

Example 2:
    Input: nums = [1,1]
    Output: [2]
 

Constraints:
    n == nums.length
    1 <= n <= 10^5
    1 <= nums[i] <= n
 

Follow up: Could you do it without extra space and in O(n) runtime? You may assume 
the returned list does not count as extra space.

"""

class Solution(object):
    def findDisappearedNumbers(self, nums):
        i = 0
        while i < len(nums):
            correct = nums[i] - 1
            if nums[i] != nums[correct]:
                nums[i], nums[correct] = nums[correct], nums[i]
            else:
                i += 1

        missing = []
        for i in range(len(nums)):
            if nums[i] != i + 1:
                missing.append(i + 1)
        return missing
