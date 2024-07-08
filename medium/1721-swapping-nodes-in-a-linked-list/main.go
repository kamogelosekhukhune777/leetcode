package main

import "fmt"

/*
You are given the head of a linked list, and an integer k.

Return the head of the linked list after swapping the values of the kth node from the beginning and
the kth node from the end (the list is 1-indexed).

Example 1:

	Input: head = [1,2,3,4,5], k = 2
	Output: [1,4,3,2,5]

Example 2:

	Input: head = [7,9,6,6,7,8,3,0,9,5], k = 5
	Output: [7,9,6,6,8,7,3,0,9,5]


Constraints:

	The number of nodes in the list is n.
	1 <= k <= n <= 105
	0 <= Node.val <= 100
*/

// ListNode defines a node in the linked list
type ListNode struct {
	Val  int
	Next *ListNode
}

// swapNodes swaps the values of the k-th node from the beginning and the k-th node from the end
func swapNodes(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}

	// Find the k-th node from the beginning
	first := head
	for i := 1; i < k; i++ {
		first = first.Next
	}

	// Use two pointers to find the k-th node from the end
	second := head
	temp := first
	for temp.Next != nil {
		temp = temp.Next
		second = second.Next
	}

	// Swap the values of the k-th nodes
	first.Val, second.Val = second.Val, first.Val

	return head
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
	k1 := 2
	fmt.Println("Original list 1:")
	printList(head1)
	modifiedHead1 := swapNodes(head1, k1)
	fmt.Println("Modified list 1:")
	printList(modifiedHead1)

	// Example 2
	head2 := &ListNode{Val: 7, Next: &ListNode{Val: 9, Next: &ListNode{Val: 6, Next: &ListNode{Val: 6, Next: &ListNode{Val: 7, Next: &ListNode{Val: 8, Next: &ListNode{Val: 3, Next: &ListNode{Val: 0, Next: &ListNode{Val: 9, Next: &ListNode{Val: 5, Next: nil}}}}}}}}}}
	k2 := 5
	fmt.Println("Original list 2:")
	printList(head2)
	modifiedHead2 := swapNodes(head2, k2)
	fmt.Println("Modified list 2:")
	printList(modifiedHead2)
}
