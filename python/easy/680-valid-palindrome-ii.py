"""
680. Valid Palindrome II

Given a string s, return true if the s can be palindrome after deleting 
at most one character from it.

 
Example 1:
    Input: s = "aba"
    Output: true

Example 2:
    Input: s = "abca"
    Output: true
    Explanation: You could delete the character 'c'.

Example 3:
    Input: s = "abc"
    Output: false
 

Constraints:
    1 <= s.length <= 10^5
    s consists of lowercase English letters.
"""

class Solution:
    def validPalindrome(s):
        def is_palindrome(left, right):
            while left < right:
                if s[left] != s[right]:
                    return False
                left += 1
                right -= 1
            return True

        left, right = 0, len(s) - 1
        while left < right:
            if s[left] == s[right]:
                left += 1
                right -= 1
            else:
                return is_palindrome(left + 1, right) or is_palindrome(left, right-1)
        return True

print("Example 1:")
s = "aba"
print(Solution.validPalindrome(s))#Output: true

print("Example 2:")
s = "abca"
print(Solution.validPalindrome(s))#Output: true

print("Example 3:")
s = "abc"
print(Solution.validPalindrome(s))#Output: false