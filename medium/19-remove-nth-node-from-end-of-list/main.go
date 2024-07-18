package main

import "fmt"

/*
Given the head of a linked list, remove the nth node from the
end of the list and return its head.

Example 1:

	Input: head = [1,2,3,4,5], n = 2
	Output: [1,2,3,5]

Example 2:

	Input: head = [1], n = 1
	Output: []

Example 3:

	Input: head = [1,2], n = 1
	Output: [1]


Constraints:

	The number of nodes in the list is sz.
	1 <= sz <= 30
	0 <= Node.val <= 100
	1 <= n <= sz


(Follow up: Could you do this in one pass?)
*/

/*
Approach
	To solve this problem using the two-pointer technique, we can follow these steps:

	1. Initialize Two Pointers: Start with two pointers, fast and slow, both pointing to the head of the list.
	2. Move Fast Pointer: Move the fast pointer n steps ahead.
	3. Move Both Pointers: Move both pointers until the fast pointer reaches the end of the list. At
	   this point, the slow pointer will be pointing to the node just before the node to be removed.
	4. Remove the Node: Adjust the next pointer of the slow node to skip the node to be removed.
	5.Edge Case Handling: Handle the edge case where the node to be removed is the head of the list.
*/

// ListNode defines a node in the linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// removeNthFromEnd removes the nth node from the end of the list.
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// Create a dummy node to simplify edge cases
	dummy := &ListNode{Next: head}
	fast, slow := dummy, dummy

	// Move the fast pointer n steps ahead
	for i := 0; i <= n; i++ {
		fast = fast.Next
	}

	// Move both pointers until the fast pointer reaches the end
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}

	// Remove the nth node from the end
	slow.Next = slow.Next.Next

	// Return the head of the modified list
	return dummy.Next
}

// Helper function to print the list
func printList(head *ListNode) {
	for head != nil {
		fmt.Print(head.Val, " ")
		head = head.Next
	}
	fmt.Println()
}

func main() {
	// Example 1
	head1 := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}
	n1 := 2
	printList(removeNthFromEnd(head1, n1)) // Output: 1 2 3 5

	// Example 2
	head2 := &ListNode{1, nil}
	n2 := 1
	printList(removeNthFromEnd(head2, n2)) // Output: (empty list)

	// Example 3
	head3 := &ListNode{1, &ListNode{2, nil}}
	n3 := 1
	printList(removeNthFromEnd(head3, n3)) // Output: 1
}
