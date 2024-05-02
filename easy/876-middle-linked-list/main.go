package main

import "fmt"

/*
Given the head of a singly linked list, return the middle node of the linked list.
If there are two middle nodes, return the second middle node.


Example 1:

Input: head = [1,2,3,4,5]
Output: [3,4,5]
Explanation: The middle node of the list is node 3.

Example 2:

Input: head = [1,2,3,4,5,6]
Output: [4,5,6]
Explanation: Since the list has two middle nodes with values 3 and 4, we return the second one.


Constraints:

The number of nodes in the list is in the range [1, 100].
1 <= Node.val <= 100
*/

//definition for singly-linked list
type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	// Example 1
	head1 := &ListNode{Val: 1}
	head1.Next = &ListNode{Val: 2}
	head1.Next.Next = &ListNode{Val: 3}
	head1.Next.Next.Next = &ListNode{Val: 4}
	head1.Next.Next.Next.Next = &ListNode{Val: 5}

	middleNode1 := middleNode(head1)
	fmt.Println("Middle Node for Example 1:", middleNode1.Val) // Output: 3

	// Example 2
	head2 := &ListNode{Val: 1}
	head2.Next = &ListNode{Val: 2}
	head2.Next.Next = &ListNode{Val: 3}
	head2.Next.Next.Next = &ListNode{Val: 4}
	head2.Next.Next.Next.Next = &ListNode{Val: 5}
	head2.Next.Next.Next.Next.Next = &ListNode{Val: 6}

	middleNode2 := middleNode(head2)
	fmt.Println("Middle Node for Example 2:", middleNode2.Val) // Output: 4
}

func middleNode(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	fast, slow := head, head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow
}

/*
func findMiddleNode(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	slow, fast := head, head

	// Move the fast pointer twice as fast as the slow pointer
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow
}*/
