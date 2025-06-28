"""
Given a string num which represents an integer, return true if 
num is a strobogrammatic number.

A strobogrammatic number is a number that looks the same when rotated 
180 degrees (looked at upside down).

Example 1:
    Input: num = "69"
    Output: true

Example 2:
    Input: num = "88"
    Output: true

Example 3:
    Input: num = "962"
    Output: false

Constraints:
    1 <= num.length <= 50
    num consists of only digits.
    num does not contain any leading zeros except for zero itself.
"""

class Solution:
    def isStrobogrammatic(num):
        rotation = {
            '0' : '0',
            '1' : '1',
            '6' : '9',
            '8' : '8',
            '9' : '6'
        }

        left, right = 0, len(num)-1
        while left <= right:
            if num[left] not in rotation or num[right] not in rotation:
                return False
            
            if rotation[num[left]] != num[right]:
                return False
            
            left += 1
            right -=1
            
        return True

num = "69"
print("example1: ",Solution.isStrobogrammatic(num))

num = "88"
print("example2: ",Solution.isStrobogrammatic(num))

num = "962"
print("example3: ",Solution.isStrobogrammatic(num))
