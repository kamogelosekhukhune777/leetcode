"""
179. Largest Number
 
Given a list of non-negative integers nums, arrange them such that they 
form the largest number and return it.

Since the result may be very large, so you need to return a string instead 
of an integer.

Example 1:
    Input: nums = [10,2]
    Output: "210"

Example 2:
    Input: nums = [3,30,34,5,9]
    Output: "9534330"
 
Constraints:
    1 <= nums.length <= 100
    0 <= nums[i] <= 10^9 
 
"""

from functools import cmp_to_key

class Solution(object):
    def largestNumber(self, nums):

        # Step 1: Convert all integers to strings
        strs = list(map(str, nums))

        # Step 2: Define custom comparator
        def compare(a, b):
            # If putting a before b gives bigger number, a should come first
            if a + b > b + a:
                return -1   # a before b
            elif a + b < b + a:
                return 1    # b before a
            else:
                return 0    # equal

        # Step 3: Sort using custom comparator
        strs.sort(key=cmp_to_key(compare))

        # Step 4: Concatenate the sorted strings
        result = "".join(strs)

        # Step 5: Handle the case of leading zeros (e.g., [0,0])
        return "0" if result[0] == "0" else result
