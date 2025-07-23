"""
1446. Consecutive Characters
 
The power of the string is the maximum length of a non-empty substring that 
contains only one unique character.

Given a string s, return the power of s.

 
Example 1:
    Input: s = "leetcode"
    Output: 2
    Explanation: The substring "ee" is of length 2 with the character 'e' only.

Example 2:
    Input: s = "abbcccddddeeeeedcba"
    Output: 5
    Explanation: The substring "eeeee" is of length 5 with the character 'e' only.
    

Constraints:
    1 <= s.length <= 500
    s consists of only lowercase English letters.
 
"""

class Solution:
    def maxPower(s):
        max_power = 1
        current_power = 1

        for i in range(1, len(s)):
            if s[i] == s[i - 1]:
                current_power += 1
                max_power = max(max_power, current_power)
            else:
                current_power = 1
        return max_power

print("Example 1:")
s = "leetcode"
print(Solution.maxPower(s))  #Output: 2
    

print("Example 2:")
s = "abbcccddddeeeeedcba"
print(Solution.maxPower(s))  #Output: 5