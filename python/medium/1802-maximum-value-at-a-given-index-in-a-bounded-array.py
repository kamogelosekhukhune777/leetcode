"""
From: https://leetcode.com/problems/maximum-value-at-a-given-index-in-a-bounded-array/description/

You are given three positive integers: n, index, and maxSum. You want to construct an array 
nums (0-indexed) that satisfies the following conditions:

    nums.length == n
    nums[i] is a positive integer where 0 <= i < n.
    abs(nums[i] - nums[i+1]) <= 1 where 0 <= i < n-1.
    The sum of all the elements of nums does not exceed maxSum.
    nums[index] is maximized.

Return nums[index] of the constructed array.

Note that abs(x) equals x if x >= 0, and -x otherwise.

Example 1:
    Input: n = 4, index = 2,  maxSum = 6
    Output: 2
    Explanation: nums = [1,2,2,1] is one array that satisfies all the conditions.
        There are no arrays that satisfy all the conditions and have nums[2] == 3, so 2 is the maximum nums[2].

Example 2:
    Input: n = 6, index = 1,  maxSum = 10
    Output: 3
 
Constraints:
    1 <= n <= maxSum <= 10^9
    0 <= index < n
"""
class Solution(object):
    def maxValue(self, n, index, maxSum):
        def calc_sum(length, peak):
            if peak >= length:
                # Full descending sequence fits
                return (peak + (peak - length + 1)) * length // 2
            else:
                # Partial sequence + 1s
                return (peak * (peak + 1)) // 2 + (length - peak)

        def is_valid(mid):
            left_len = index
            right_len = n - index - 1
            total = calc_sum(left_len, mid - 1) + calc_sum(right_len, mid - 1) + mid
            return total <= maxSum

        low, high = 1, maxSum
        res = 1

        while low <= high:
            mid = (low + high) // 2
            if is_valid(mid):
                res = mid
                low = mid + 1
            else:
                high = mid - 1

        return res

