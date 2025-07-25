"""
922. Sort Array By Parity II

Given an array of integers nums, half of the integers in nums are odd, and the other half are even.

Sort the array so that whenever nums[i] is odd, i is odd, and whenever nums[i] is even, i is even.

Return any answer array that satisfies this condition.

 
Example 1:
    Input: nums = [4,2,5,7]
    Output: [4,5,2,7]
    Explanation: [4,7,2,5], [2,5,4,7], [2,7,4,5] would also have been accepted.

Example 2:
    Input: nums = [2,3]
    Output: [2,3]
 

Constraints:
    2 <= nums.length <= 2 * 10^4
    nums.length is even.
    Half of the integers in nums are even.
    0 <= nums[i] <= 1000
 

Follow Up: Could you solve it in-place?
"""

class Solution:
    def sortArrayByParityII(nums):
        i, j = 0, 1
        while i < len(nums) and j < len(nums):
            if nums[i] % 2 == 0:
                i += 2
            elif nums[j] % 2 == 1:
                j += 2
            else:
                # nums[i] is odd, nums[j] is even => swap
                nums[i], nums[j] = nums[j], nums[i]
                i += 2
                j += 2
        return nums

print("Example: 1")
nums = [4,2,5,7]
print(Solution.sortArrayByParityII(nums))

print("Example: 2")
nums = [2,3]
print(Solution.sortArrayByParityII(nums))