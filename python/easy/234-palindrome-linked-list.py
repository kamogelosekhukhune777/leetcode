"""
234. Palindrome Linked List

Given the head of a singly linked list, return true if it is a palindrome or false otherwise.

 
Example 1:
    Input: head = [1,2,2,1]
    Output: true

Example 2:
    Input: head = [1,2]
    Output: false
 

Constraints:
    The number of nodes in the list is in the range [1, 105].
    0 <= Node.val <= 9
 

Follow up: Could you do it in O(n) time and O(1) space?
"""

# Definition for singly-linked list.
class ListNode(object):
    def __init__(self, val=0, next=None):
         self.val = val
         self.next = next

class Solution(object):
    def isPalindrome(self, head):
        if not head or not head.next:
            return True
        
        # 1. Find middle (slow/fast)
        slow, fast = head, head
        while fast and fast.next:
            slow = slow.next
            fast = fast.next.next
        
        # 2. Reverse second half
        prev, curr = None, slow
        while curr:
            nxt = curr.next
            curr.next = prev
            prev, curr = curr, nxt
        
        # 3. Compare halves
        left, right = head, prev
        while right:  # only need to check second half
            if left.val != right.val:
                return False
            left, right = left.next, right.next
        return True
