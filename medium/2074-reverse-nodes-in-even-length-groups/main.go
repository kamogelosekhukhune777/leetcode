package main

import (
	"fmt"
)

/*
You are given the head of a linked list.

The nodes in the linked list are sequentially assigned to non-empty groups whose lengths form the sequence of
the natural numbers (1, 2, 3, 4, ...). The length of a group is the number of nodes assigned to it. In other words,

	The 1st node is assigned to the first group.
	The 2nd and the 3rd nodes are assigned to the second group.
	The 4th, 5th, and 6th nodes are assigned to the third group, and so on.

Note that the length of the last group may be less than or equal to 1 + the length of the second to last group.

Reverse the nodes in each group with an even length, and return the head of the modified linked list.



Example 1:

	Input: head = [5,2,6,3,9,1,7,3,8,4]
	Output: [5,6,2,3,9,1,4,8,3,7]
	Explanation:
		- The length of the first group is 1, which is odd, hence no reversal occurs.
		- The length of the second group is 2, which is even, hence the nodes are reversed.
		- The length of the third group is 3, which is odd, hence no reversal occurs.
		- The length of the last group is 4, which is even, hence the nodes are reversed.

Example 2:

	Input: head = [1,1,0,6]
	Output: [1,0,1,6]
	Explanation:
		- The length of the first group is 1. No reversal occurs.
		- The length of the second group is 2. The nodes are reversed.
		- The length of the last group is 1. No reversal occurs.
Example 3:

	Input: head = [1,1,0,6,5]
	Output: [1,0,1,5,6]
	Explanation:
		- The length of the first group is 1. No reversal occurs.
		- The length of the second group is 2. The nodes are reversed.
		- The length of the last group is 2. The nodes are reversed.


Constraints:

	The number of nodes in the list is in the range [1, 105].
	0 <= Node.val <= 10^5
*/

// ListNode defines a node in the linked list
type ListNode struct {
	Val  int
	Next *ListNode
}

// reverse function reverses a part of the linked list from start to end and returns the new head and tail
func reverse(start, end *ListNode) (*ListNode, *ListNode) {
	ptr := start
	var prev *ListNode = nil

	for ptr != nil {
		next := ptr.Next
		ptr.Next = prev
		prev = ptr
		ptr = next
	}

	return end, start
}

// reverseEvenLengthGroups function reverses nodes in even-length groups in a linked list
func reverseEvenLengthGroups(head *ListNode) *ListNode {
	// If no elements
	if head == nil {
		return nil
	}

	// Start from the second node, we know we are not going to reverse size 1 anyway
	groupSize := 2

	start, end := head.Next, head.Next
	var prev *ListNode = head

	for start != nil {
		size := 1
		for end.Next != nil && size < groupSize {
			end = end.Next
			size++
		}

		// Before we reverse let's just keep track of what's going to be our next
		next := end.Next

		// We need to reverse here
		if size%2 == 0 {
			// Cut it off from the future list
			end.Next = nil

			// Reverse the pointers between start and end and swap their values
			start, end = reverse(start, end)

			// Update the last node prev to this group with updated next value
			prev.Next = start
		}

		// Reconnect link in case we cut it off in size % 2
		end.Next = next

		// Now prev is this new end
		prev = end

		// Move pointers to next group
		start, end = next, next

		// Increase our expected group size by 1
		groupSize++
	}

	return head
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
	head1 := createLinkedList([]int{5, 2, 6, 3, 9, 1, 7, 3, 8, 4})
	head2 := createLinkedList([]int{1, 1, 0, 6})
	head3 := createLinkedList([]int{1, 1, 0, 6, 5})

	head1 = reverseEvenLengthGroups(head1)
	head2 = reverseEvenLengthGroups(head2)
	head3 = reverseEvenLengthGroups(head3)

	fmt.Println("Output for head1:")
	printLinkedList(head1)

	fmt.Println("Output for head2:")
	printLinkedList(head2)

	fmt.Println("Output for head3:")
	printLinkedList(head3)
}
