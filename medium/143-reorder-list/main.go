package main

import "fmt"

/*
You are given the head of a singly linked-list. The list can be represented as:

	L0 → L1 → … → Ln - 1 → Ln
Reorder the list to be on the following form:

	L0 → Ln → L1 → Ln - 1 → L2 → Ln - 2 → …
You may not modify the values in the list's nodes. Only nodes themselves may be changed.


Example 1:

	Input: head = [1,2,3,4]
	Output: [1,4,2,3]

Example 2:

	Input: head = [1,2,3,4,5]
	Output: [1,5,2,4,3]


Constraints:

	The number of nodes in the list is in the range [1, 5 * 104].
	1 <= Node.val <= 1000
*/

// ListNode defines a node in the linked list
type ListNode struct {
	Val  int
	Next *ListNode
}

// reorderList reorders the list as specified
func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}

	// Step 1: Find the middle of the list
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	// Step 2: Reverse the second half of the list
	secondHalf := reverseList(slow.Next)
	slow.Next = nil // Split the list into two halves

	// Step 3: Merge the two halves
	firstHalf := head
	mergeLists(firstHalf, secondHalf)
}

// reverseList reverses the linked list and returns the new head
func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
}

// mergeLists merges two lists alternately
func mergeLists(l1, l2 *ListNode) {
	for l1 != nil && l2 != nil {
		l1Next, l2Next := l1.Next, l2.Next
		l1.Next = l2
		if l1Next == nil {
			break
		}
		l2.Next = l1Next
		l1 = l1Next
		l2 = l2Next
	}
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
	head1 := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: nil}}}}
	fmt.Println("Original list 1:")
	printList(head1)
	reorderList(head1)
	fmt.Println("Reordered list 1:")
	printList(head1)

	// Example 2
	head2 := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: nil}}}}}
	fmt.Println("Original list 2:")
	printList(head2)
	reorderList(head2)
	fmt.Println("Reordered list 2:")
	printList(head2)
}
