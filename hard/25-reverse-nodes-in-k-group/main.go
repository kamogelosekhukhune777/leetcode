package main

import "fmt"

/*
Given the head of a linked list, reverse the nodes of the list k at a time, and return the modified list.

k is a positive integer and is less than or equal to the length of the linked list. If the number
of nodes is not a multiple of k then left-out nodes, in the end, should remain as it is.

You may not alter the values in the list's nodes, only nodes themselves may be changed.



Example 1:

	Input: head = [1,2,3,4,5], k = 2
	Output: [2,1,4,3,5]

Example 2:

	Input: head = [1,2,3,4,5], k = 3
	Output: [3,2,1,4,5]


Constraints:

	The number of nodes in the list is n.
	1 <= k <= n <= 5000
	0 <= Node.val <= 1000


Follow-up: Can you solve the problem in O(1) extra memory space?
*/

// ListNode defines a node in the linked list
type ListNode struct {
	Val  int
	Next *ListNode
}

// reverseKGroup reverses nodes in k-sized groups
func reverseKGroup(head *ListNode, k int) *ListNode {
	dummy := &ListNode{Next: head}
	prevGroupEnd := dummy

	for {
		// Check if there are k nodes left to reverse
		kthNode := getKthNode(prevGroupEnd, k)
		if kthNode == nil {
			break
		}

		// Pointers to the current group
		//sgroupStart := prevGroupEnd.Next
		groupEnd := kthNode.Next

		// Reverse the nodes in the current group
		prev, curr := kthNode.Next, prevGroupEnd.Next
		for curr != groupEnd {
			next := curr.Next
			curr.Next = prev
			prev = curr
			curr = next
		}

		// Connect the previous group to the reversed group
		temp := prevGroupEnd.Next
		prevGroupEnd.Next = kthNode
		prevGroupEnd = temp
	}

	return dummy.Next
}

// getKthNode returns the k-th node from the given node
func getKthNode(start *ListNode, k int) *ListNode {
	current := start
	for i := 0; i < k; i++ {
		if current == nil {
			return nil
		}
		current = current.Next
	}
	return current
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
	// Create a linked list: 1 -> 2 -> 3 -> 4 -> 5 -> nil
	head := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: nil}}}}}

	fmt.Println("Original list:")
	printList(head)

	k := 2
	modifiedHead := reverseKGroup(head, k)

	fmt.Println("Modified list (k=2):")
	printList(modifiedHead)

	// Reset the list for another example
	head = &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: nil}}}}}
	k = 3
	modifiedHead = reverseKGroup(head, k)

	fmt.Println("Modified list (k=3):")
	printList(modifiedHead)
}
