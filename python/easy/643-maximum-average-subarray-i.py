"""
643. Maximum Average Subarray I
Easy
Topics
premium lock icon
Companies
You are given an integer array nums consisting of n elements, and an integer k.

Find a contiguous subarray whose length is equal to k that has the maximum average 
value and return this value. Any answer with a calculation error less 
than 10-5 will be accepted.


Example 1:
    Input: nums = [1,12,-5,-6,50,3], k = 4
    Output: 12.75000
    Explanation: Maximum average is (12 - 5 - 6 + 50) / 4 = 51 / 4 = 12.75

Example 2:
    Input: nums = [5], k = 1
    Output: 5.00000

 
Constraints:
    n == nums.length
    1 <= k <= n <= 10^5
    -10^^4 <= nums[i] <= 10^4
"""

class Solution:
    def findMaxAverage(nums, k):
        window_sum = sum(nums[:k])
        max_sum = window_sum
        
        for i in range(k, len(nums)):
            window_sum += nums[i] - nums[i-k]
            max_sum = max(max_sum, window_sum)
        
        return float(max_sum)/k

#Example 1:
print("Example 1:")
nums = [1,12,-5,-6,50,3]
k = 4
print(Solution.findMaxAverage(nums,k))

#Example 2:
print("Example 1:")
nums = [5]
k = 1
print(Solution.findMaxAverage(nums,k))
