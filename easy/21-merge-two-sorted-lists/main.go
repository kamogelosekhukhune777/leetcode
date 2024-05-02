package main

import "fmt"

/*
You are given the heads of two sorted linked lists list1 and list2.

Merge the two lists into one sorted list. The list should be made by
splicing together the nodes of the first two lists.

Return the head of the merged linked list.

Example 1:

Input: list1 = [1,2,4], list2 = [1,3,4]
Output: [1,1,2,3,4,4]

Example 2:

Input: list1 = [], list2 = []
Output: []

Example 3:

Input: list1 = [], list2 = [0]
Output: [0]

Constraints:
   The number of nodes in both lists is in the range [0, 50].
   -100 <= Node.val <= 100
   Both list1 and list2 are sorted in non-decreasing order.

*/

// ListNode represents a node in the linked list
type ListNode struct {
	Val  int
	Next *ListNode
}

// mergeTwoLists merges two sorted linked lists into one sorted list
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	// Create a dummy node to start the merged list
	dummy := &ListNode{}
	current := dummy

	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			current.Next = l1
			l1 = l1.Next
		} else {
			current.Next = l2
			l2 = l2.Next
		}
		current = current.Next
	}

	// Attach remaining nodes from l1 or l2 if there are any
	if l1 != nil {
		current.Next = l1
	} else {
		current.Next = l2
	}

	return dummy.Next
}

// printLinkedList prints all the node values of a linked list
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
	list1 := &ListNode{Val: 1}
	list1.Next = &ListNode{Val: 2}
	list1.Next.Next = &ListNode{Val: 4}

	list2 := &ListNode{Val: 1}
	list2.Next = &ListNode{Val: 3}
	list2.Next.Next = &ListNode{Val: 4}

	mergedList1 := mergeTwoLists(list1, list2)
	printLinkedList(mergedList1) // Output: 1 1 2 3 4 4

	// Example 2
	list3 := &ListNode{}

	list4 := &ListNode{}

	mergedList2 := mergeTwoLists(list3, list4)
	printLinkedList(mergedList2) // Output: (empty list)

	// Example 3
	list5 := &ListNode{}

	list6 := &ListNode{Val: 0}

	mergedList3 := mergeTwoLists(list5, list6)
	printLinkedList(mergedList3)
}
