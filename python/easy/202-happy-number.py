"""
Write an algorithm to determine if a number n is happy.

A happy number is a number defined by the following process:

    -Starting with any positive integer, replace the number by the sum of 
     the squares of its digits.
    -Repeat the process until the number equals 1 (where it will stay), or 
     it loops endlessly in a cycle which does not include 1.
    -Those numbers for which this process ends in 1 are happy.
Return true if n is a happy number, and false if not.

Example 1:
    Input: n = 19
    Output: true
    Explanation:
        12 + 92 = 82
        82 + 22 = 68
        62 + 82 = 100
        12 + 02 + 02 = 1

Example 2:
    Input: n = 2
    Output: false

Constraints:
    1 <= n <= 2^31 - 1

"""

class Solution1(object):
    def isHappy(self, n):
        def get_next(n):
            total = 0
            while n > 0:
                digit = n % 10
                total += digit * digit
                n //= 10
            return total


        slow = n
        fast = get_next(n)
        
        while fast != 1 and slow != fast:
            slow = get_next(slow)
            fast = get_next(get_next(fast))
        
        return fast == 1

class Solution2(object):
    def isHappy(self, n):
        seen = set()
        while n != 1:
            if n in seen:
                return False
            seen.add(n)
            n = sum(int(digit)**2 for digit in str(n))
        return True
