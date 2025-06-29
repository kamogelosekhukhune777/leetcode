"""
You are given an integer num. You can swap two digits at most 
once to get the maximum valued number.

Return the maximum valued number you can get.


Example 1:
    Input: num = 2736
    Output: 7236
    Explanation: Swap the number 2 and the number 7.

Example 2:
    Input: num = 9973
    Output: 9973
    Explanation: No swap.
 
Constraints:
    0 <= num <= 10^8
"""

class Solution(object):
    def maximumSwap(self, num):
        digits = list(str(num))
        last = {int(d): i for i, d in enumerate(digits)}  # last index of each digit

        for i, d in enumerate(digits):
            for k in range(9, int(d), -1):  # check for bigger digits from 9 down to d+1
                if last.get(k, -1) > i:
                    # swap and return
                    digits[i], digits[last[k]] = digits[last[k]], digits[i]
                    return int("".join(digits))
        
        return num