package main

import (
	"fmt"
)

/* https://leetcode.com/problems/swap-nodes-in-pairs/

Given a linked list, swap every two adjacent nodes and return its head. You must solve
the problem without modifying the values in the list's nodes (i.e., only nodes themselves may be changed.)


Example 1:

	Input: head = [1,2,3,4]
	Output: [2,1,4,3]

Example 2:

	Input: head = []
	Output: []

Example 3:

	Input: head = [1]
	Output: [1]

Constraints:

	The number of nodes in the list is in the range [0, 100].
	0 <= Node.val <= 100
*/

// ListNode defines a node in the linked list
type ListNode struct {
	Val  int
	Next *ListNode
}

// Function to swap every two adjacent nodes in the linked list
func swapPairs(head *ListNode) *ListNode {
	// Initialize a dummy node to handle edge cases cleanly
	dummy := &ListNode{Next: head}
	// prev is a pointer to the last node before the current pair
	prev := dummy

	// Iterate through the list
	for head != nil && head.Next != nil {
		// Identify the two nodes to be swapped
		first := head
		second := head.Next

		// Swap the nodes by adjusting the pointers
		first.Next = second.Next
		second.Next = first
		prev.Next = second

		// Move prev to the end of the swapped pair
		prev = first
		// Move head to the next pair of nodes
		head = first.Next
	}

	return dummy.Next
}

// Helper function to create a linked list from a slice
func createLinkedList(nums []int) *ListNode {
	dummy := &ListNode{}
	curr := dummy
	for _, num := range nums {
		curr.Next = &ListNode{Val: num}
		curr = curr.Next
	}
	return dummy.Next
}

// Helper function to print the linked list
func printLinkedList(head *ListNode) {
	for head != nil {
		fmt.Print(head.Val, " ")
		head = head.Next
	}
	fmt.Println()
}

func main() {
	// Test cases
	head1 := createLinkedList([]int{1, 2, 3, 4})
	head2 := createLinkedList([]int{})
	head3 := createLinkedList([]int{1})

	head1 = swapPairs(head1)
	head2 = swapPairs(head2)
	head3 = swapPairs(head3)

	fmt.Println("Output for head1:")
	printLinkedList(head1) // Output: [2, 1, 4, 3]
	fmt.Println("Output for head2:")
	printLinkedList(head2) // Output: []
	fmt.Println("Output for head3:")
	printLinkedList(head3) // Output: [1]
}
