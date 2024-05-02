package main

import "fmt"

//PALINDROME = a word, phrase, or sequence that reads the same backwards as forwards
/*
234. Palindrome Linked List

Given the head of a singly linked list, return true if it is a palindromeor false otherwise.

Example 1:

    1 -> 2 -> 2 -> 1

Input: head = [1,2,2,1]
Output: true

Example 2:

    1 -> 2

Input: head = [1,2]
Output: false


Constraints:

The number of nodes in the list is in the range [1, 105].
0 <= Node.val <= 9
*/

//definition for singly-linked list
type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	//test Cases
	//Example 1
	head1 := &ListNode{Val: 1}
	head1.Next = &ListNode{Val: 2}
	head1.Next.Next = &ListNode{Val: 2}
	head1.Next.Next.Next = &ListNode{Val: 1}
	fmt.Println("is Palindrome", isPalindrome(head1))

	//Example 2
	head2 := &ListNode{Val: 1}
	head2.Next = &ListNode{Val: 2}
	fmt.Println("is Palindrome:", isPalindrome(head2))
}

func isPalindrome(head *ListNode) bool {
	//edge case
	if head == nil || head.Next == nil {
		return true
	}
	//find the middle node
	var slow, fast *ListNode = head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	//reverse the second half of list
	secondHalf := reverse(slow)

	//compare first and second half
	for secondHalf != nil {
		if head.Val != secondHalf.Val {
			return false
		}
		head = head.Next
		secondHalf = secondHalf.Next
	}

	return true
}

func reverse(head *ListNode) *ListNode {
	var prev, current *ListNode = nil, head

	for current != nil {
		next := current.Next
		current.Next = prev
		prev = current
		current = next
	}

	return prev
}
