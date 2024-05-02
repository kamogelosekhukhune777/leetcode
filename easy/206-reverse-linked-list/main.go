package main

import "fmt"

/*
Given the head of a singly linked list, reverse the list, and return the reversed list.

Example 1:

Input: head = [1,2,3,4,5]
Output: [5,4,3,2,1]

Example 2:

Input: head = [1,2]
Output: [2,1]

Example 3:

Input: head = []
Output: []


Constraints:
-    The number of nodes in the list is the range [0, 5000].
-    -5000 <= Node.val <= 5000

*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var prev, current *ListNode = nil, head

	for current != nil {
		next := current.Next

		current.Next = prev
		prev = current
		current = next
	}
	return prev
}

func main() {
	//Example 1
	List1 := &ListNode{Val: 1}
	List1.Next = &ListNode{Val: 2}
	List1.Next.Next = &ListNode{Val: 3}
	List1.Next.Next.Next = &ListNode{Val: 4}
	List1.Next.Next.Next.Next = &ListNode{Val: 5}

	reverseList1 := reverseList(List1)
	fmt.Print("reverse list1: ")
	printLinkedList(reverseList1)

	//Example 2
	List2 := &ListNode{Val: 1}
	List2.Next = &ListNode{Val: 2}

	reverseList2 := reverseList(List2)
	fmt.Print("reverse list2: ")
	printLinkedList(reverseList2)

	//Example 3
	List3 := &ListNode{}

	reverseList3 := reverseList(List3)
	fmt.Print("reverse list3: ")
	printLinkedList(reverseList3)
}

// printLinkedList prints all the node values of a linked list
//helper func
func printLinkedList(head *ListNode) {
	current := head
	for current != nil {
		fmt.Print(current.Val, " ")
		current = current.Next
	}
	fmt.Println()
}
