package main

import "fmt"

/*
Given the head of a singly linked list and two integers left and right where left <= right, reverse
the nodes of the list from position left to position right, and return the reversed list.


Example 1:

	Input: head = [1,2,3,4,5], left = 2, right = 4
	Output: [1,4,3,2,5]

Example 2:

	Input: head = [5], left = 1, right = 1
	Output: [5]


Constraints:

	The number of nodes in the list is n.
	1 <= n <= 500
	-500 <= Node.val <= 500
	1 <= left <= right <= n


Follow up: Could you do it in one pass?
*/

// ListNode defines a node in the linked list
type ListNode struct {
	Val  int
	Next *ListNode
}

// reverseBetween reverses the nodes from position left to right
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if head == nil || left == right {
		return head
	}

	dummy := &ListNode{Next: head}
	prev := dummy

	// Move prev to the node just before the left position
	for i := 1; i < left; i++ {
		prev = prev.Next
	}

	// Start is the first node of the sublist to be reversed
	start := prev.Next
	then := start.Next

	// Reverse the sublist from left to right
	for i := 0; i < right-left; i++ {
		start.Next = then.Next
		then.Next = prev.Next
		prev.Next = then
		then = start.Next
	}

	return dummy.Next
}

// printList prints the linked list
func printList(head *ListNode) {
	current := head
	for current != nil {
		fmt.Printf("%d -> ", current.Val)
		current = current.Next
	}
	fmt.Println("nil")
}

func main() {
	// Example 1
	head1 := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: nil}}}}}
	left1, right1 := 2, 4

	fmt.Println("Original list 1:")
	printList(head1)

	modifiedHead1 := reverseBetween(head1, left1, right1)

	fmt.Println("Modified list 1:")
	printList(modifiedHead1)

	// Example 2
	head2 := &ListNode{Val: 5}
	left2, right2 := 1, 1

	fmt.Println("Original list 2:")
	printList(head2)

	modifiedHead2 := reverseBetween(head2, left2, right2)

	fmt.Println("Modified list 2:")
	printList(modifiedHead2)
}
