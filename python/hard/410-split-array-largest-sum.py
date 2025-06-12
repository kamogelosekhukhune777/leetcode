"""
Given an integer array nums and an integer k, split nums into k non-empty subarrays 
such that the largest sum of any subarray is minimized.

Return the minimized largest sum of the split.

A subarray is a contiguous part of the array.


Example 1:
    Input: nums = [7,2,5,10,8], k = 2
    Output: 18
    Explanation: There are four ways to split nums into two subarrays.
        The best way is to split it into [7,2,5] and [10,8], where the largest sum among 
        the two subarrays is only 18.

Example 2:
    Input: nums = [1,2,3,4,5], k = 2
    Output: 9
    Explanation: There are four ways to split nums into two subarrays.
        The best way is to split it into [1,2,3] and [4,5], where the largest sum among 
        the two subarrays is only 9.
 
Constraints:
    1 <= nums.length <= 1000
    0 <= nums[i] <= 10^6
    1 <= k <= min(50, nums.length)
"""

class Solution(object):
    def splitArray(self, nums, k):
        def canSplit(mid):
            count = 1
            current_sum = 0
            for num in nums:
                if current_sum + num > mid:
                    count += 1
                    current_sum = num
                else:
                    current_sum += num
            return count <= k

        low, high = max(nums), sum(nums)
        result = high

        while low <= high:
            mid = (low + high) // 2
            if canSplit(mid):
                result = mid 
                high = mid - 1
            else:
                low = mid + 1

        return result

        