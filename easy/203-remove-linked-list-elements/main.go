package main

import "fmt"

/*
Given the head of a linked list and an integer val, remove all the nodes of
the linked list that has Node.val == val, and return the new head.

Example 1:

Input: head = [1,2,6,3,4,5,6], val = 6
Output: [1,2,3,4,5]

Example 2:

Input: head = [], val = 1
Output: []

Example 3:

Input: head = [7,7,7,7], val = 7
Output: []

Constraints:
-  The number of nodes in the list is in the range [0, 104].
-  1 <= Node.val <= 50
-  0 <= val <= 50
*/

// ListNode represents a node in the linked list
type ListNode struct {
	Val  int
	Next *ListNode
}

// removeElements removes all nodes from the linked list that has Node.Val == val
func removeElements(head *ListNode, val int) *ListNode {
	// Create a dummy node to handle cases where the head itself needs to be removed
	dummy := &ListNode{Next: head}
	slow, fast := dummy, head

	for fast != nil {
		if fast.Val == val {
			slow.Next = fast.Next
			fast = slow.Next
		} else {
			slow = slow.Next
			fast = fast.Next
		}
	}

	return dummy.Next
}

// printLinkedList prints all the node values of a linked list
//helper function not neccessary
func printLinkedList(head *ListNode) {
	current := head
	for current != nil {
		fmt.Print(current.Val, " ")
		current = current.Next
	}
	fmt.Println()
}

func main() {
	// Example 1
	head1 := &ListNode{Val: 1}
	head1.Next = &ListNode{Val: 2}
	head1.Next.Next = &ListNode{Val: 6}
	head1.Next.Next.Next = &ListNode{Val: 3}
	head1.Next.Next.Next.Next = &ListNode{Val: 4}
	head1.Next.Next.Next.Next.Next = &ListNode{Val: 5}
	head1.Next.Next.Next.Next.Next.Next = &ListNode{Val: 6}

	val1 := 6
	newHead1 := removeElements(head1, val1)
	printLinkedList(newHead1) // Output: 1 2 3 4 5

	// Example 2
	head2 := &ListNode{}

	val2 := 1
	newHead2 := removeElements(head2, val2)
	printLinkedList(newHead2) // Output: (empty list)

	// Example 3

	head3 := &ListNode{Val: 7}
	head3.Next = &ListNode{Val: 7}
	head3.Next.Next = &ListNode{Val: 7}
	head3.Next.Next.Next = &ListNode{Val: 7}

	val3 := 7
	newHead3 := removeElements(head3, val3)
	printLinkedList(newHead3) // Output: (empty list)
}
