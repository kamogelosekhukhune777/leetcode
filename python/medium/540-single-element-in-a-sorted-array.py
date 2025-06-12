"""
You are given a sorted array consisting of only integers where every element 
appears exactly twice, except for one element which appears exactly once.

Return the single element that appears only once.

Your solution must run in O(log n) time and O(1) space.

Example 1:
    Input: nums = [1,1,2,3,3,4,4,8,8]
    Output: 2

Example 2:
    Input: nums = [3,3,7,7,10,11,11]
    Output: 10
 

Constraints:
    1 <= nums.length <= 10^5
    0 <= nums[i] <= 10^5
"""

class Solution(object):
    def singleNonDuplicate(self, nums):
        low, high = 0, len(nums) - 1

        while low < high:
            mid = (low + high) // 2
            if mid % 2 == 1:
                mid -= 1

            if nums[mid] == nums[mid + 1]:
                low = mid + 2
            else:
                high = mid

        return nums[low]
