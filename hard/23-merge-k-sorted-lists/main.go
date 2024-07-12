package main

import (
	"container/heap"
	"fmt"
)

/*
You are given an array of k linked-lists lists, each linked-list is sorted in ascending order.

Merge all the linked-lists into one sorted linked-list and return it.


Example 1:

	Input: lists = [[1,4,5],[1,3,4],[2,6]]
	Output: [1,1,2,3,4,4,5,6]
	Explanation: The linked-lists are:
		[
		1->4->5,
		1->3->4,
		2->6
		]
merging them into one sorted list:
1->1->2->3->4->4->5->6

Example 2:

	Input: lists = []
	Output: []

Example 3:

	Input: lists = [[]]
	Output: []


Constraints:

	k == lists.length
	0 <= k <= 10^4
	0 <= lists[i].length <= 500
	-10^4 <= lists[i][j] <= 10^4
	lists[i] is sorted in ascending order.
The sum of lists[i].length will not exceed 10^4.
*/

// ListNode defines a singly linked list node
type ListNode struct {
	Val  int
	Next *ListNode
}

// Pair is a struct that stores a ListNode and its value
type Pair struct {
	node *ListNode
	val  int
}

// PriorityQueue implements heap.Interface and holds Pairs
type PriorityQueue []*Pair

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].val < pq[j].val
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(*Pair))
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

func mergeKLists(lists []*ListNode) *ListNode {
	pq := &PriorityQueue{}
	heap.Init(pq)

	// Initialize the heap with the head of each linked list
	for _, list := range lists {
		if list != nil {
			heap.Push(pq, &Pair{node: list, val: list.Val})
		}
	}

	// Dummy node to start the merged list
	dummy := &ListNode{}
	current := dummy

	// Extract the smallest node and add to the merged list
	for pq.Len() > 0 {
		p := heap.Pop(pq).(*Pair)
		current.Next = p.node
		current = current.Next
		if p.node.Next != nil {
			heap.Push(pq, &Pair{node: p.node.Next, val: p.node.Next.Val})
		}
	}

	return dummy.Next
}

func printList(head *ListNode) {
	for head != nil {
		fmt.Print(head.Val, " ")
		head = head.Next
	}
	fmt.Println()
}

func main() {
	// Example 1
	list1 := &ListNode{Val: 1, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}
	list2 := &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}}
	list3 := &ListNode{Val: 2, Next: &ListNode{Val: 6}}
	lists := []*ListNode{list1, list2, list3}
	mergedList := mergeKLists(lists)
	printList(mergedList) // Output: 1 1 2 3 4 4 5 6

	// Example 2
	lists = []*ListNode{}
	mergedList = mergeKLists(lists)
	printList(mergedList) // Output: (Empty list)

	// Example 3
	lists = []*ListNode{nil}
	mergedList = mergeKLists(lists)
	printList(mergedList) // Output: (Empty list)
}
