"""
Given a circular linked list list of positive integers, your task is to split it 
into 2 circular linked lists so that the first one contains the first half of the 
nodes in list (exactly ceil(list.length / 2) nodes) in the same order they appeared 
in list, and the second one contains the rest of the nodes in list in the same 
order they appeared in list.

Return an array answer of length 2 in which the first element is a circular linked 
list representing the first half and the second element is a circular linked list 
representing the second half.

A circular linked list is a normal linked list with the only difference being that 
the last node's next node, is the first node.
 

Example 1:
    Input: nums = [1,5,7]
    Output: [[1,5],[7]]
    Explanation: The initial list has 3 nodes so the first half would be the 
        first 2 elements since ceil(3 / 2) = 2 and the rest which is 1 node 
        is in the second half.

Example 2:
    Input: nums = [2,6,1,5]
    Output: [[2,6],[1,5]]
    Explanation: The initial list has 4 nodes so the first half would be the 
        first 2 elements since ceil(4 / 2) = 2 and the rest which is 2 nodes 
        are in the second half.
 

Constraints:
    -The number of nodes in list is in the range [2, 10^5]
    - 0 <= Node.val <= 10^9
    -LastNode.next = FirstNode where LastNode is the last node of the list 
     and FirstNode is the first one
"""

class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next

#(Fast & Slow Pointer - O(1) space)
class Solution1:
    def splitCircularList(head):
        if not head or not head.next or head.next == head:
            return [head, None]

        slow = head
        fast = head

        # Step 1: Find the midpoint using fast and slow pointers
        while fast.next != head and fast.next.next != head:
            slow = slow.next
            fast = fast.next.next

        first_half_head = head
        second_half_head = slow.next

        # Step 2: Make the first half circular
        slow.next = first_half_head

        # Step 3: Make the second half circular
        curr = second_half_head
        while curr.next != head:
            curr = curr.next
        curr.next = second_half_head

        return [first_half_head, second_half_head]

#=============================================================================

class Solution2:
    def build_circular_list(values):
        if not values:
            return None
        head = ListNode(values[0])
        curr = head
        for v in values[1:]:
            curr.next = ListNode(v)
            curr = curr.next
        curr.next = head  # make it circular
        return head

def to_array(head, limit):
    # Convert a circular linked list to array with length = limit
    result = []
    curr = head
    for _ in range(limit):
        result.append(curr.val)
        curr = curr.next
    return result

