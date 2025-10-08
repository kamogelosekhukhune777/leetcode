"""
2099. Find Subsequence of Length K With the Largest Sum

You are given an integer array nums and an integer k. You want to find a subsequence of nums of length k 
that has the largest sum.

Return any such subsequence as an integer array of length k.

A subsequence is an array that can be derived from another array by deleting some or no elements without 
changing the order of the remaining elements.


Example 1:
    Input: nums = [2,1,3,3], k = 2
    Output: [3,3]
    Explanation:
        The subsequence has the largest sum of 3 + 3 = 6.

Example 2:
    Input: nums = [-1,-2,3,4], k = 3
    Output: [-1,3,4]
    Explanation: 
        The subsequence has the largest sum of -1 + 3 + 4 = 6.

Example 3:
    Input: nums = [3,4,3,3], k = 2
    Output: [3,4]
    Explanation:
        The subsequence has the largest sum of 3 + 4 = 7. 
        Another possible subsequence is [4, 3].

Constraints:
    1 <= nums.length <= 1000
    -10^5 <= nums[i] <= 10^5
    1 <= k <= nums.length
"""

import heapq

class Solution(object):
    def maxSubsequence(self, nums, k):
        # Step 1: Build list of (num, index)
        heap = []
        for i, num in enumerate(nums):
            heapq.heappush(heap, (num, i))
            if len(heap) > k:
                heapq.heappop(heap)
        
        # Step 2: Sort top k elements by index
        top_k = sorted(heap, key=lambda x: x[1])
        
        # Step 3: Return only values
        return [num for num, _ in top_k]
