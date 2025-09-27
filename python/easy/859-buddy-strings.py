"""
859. Buddy Strings

Given two strings s and goal, return true if you can swap two letters in s so the result is equal to goal, 
otherwise, return false.

Swapping letters is defined as taking two indices i and j (0-indexed) such that i != j and swapping the 
characters at s[i] and s[j].

For example, swapping at indices 0 and 2 in "abcd" results in "cbad".


Example 1:
    Input: s = "ab", goal = "ba"
    Output: true
    Explanation: You can swap s[0] = 'a' and s[1] = 'b' to get "ba", which is equal to goal.

Example 2:
    Input: s = "ab", goal = "ab"
    Output: false
    Explanation: The only letters you can swap are s[0] = 'a' and s[1] = 'b', which results in "ba" != goal.

Example 3:
    Input: s = "aa", goal = "aa"
    Output: true
    Explanation: You can swap s[0] = 'a' and s[1] = 'a' to get "aa", which is equal to goal.

Constraints:
    1 <= s.length, goal.length <= 2 * 10^4
    s and goal consist of lowercase letters.
"""

class Solution(object):
    def buddyStrings(self, s, goal):
        if len(s) != len(goal):
            return False
        
        if s == goal:
            return len(set(s)) < len(s)  # has duplicate
        
        mismatches = []
        for i in range(len(s)):
            if s[i] != goal[i]:
                mismatches.append(i)
        
        if len(mismatches) != 2:
            return False
        
        i, j = mismatches
        return s[i] == goal[j] and s[j] == goal[i]