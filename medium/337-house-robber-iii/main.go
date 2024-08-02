package main

import "fmt"

/*
The thief has found himself a new place for his thievery again. There is only
one entrance to this area, called root.

Besides the root, each house has one and only one parent house. After a tour, the
smart thief realized that all houses in this place form a binary tree. It will automatically
contact the police if two directly-linked houses were broken into on the same night.

Given the root of the binary tree, return the maximum amount of money the thief can
rob without alerting the police.


Example 1:

	Input: root = [3,2,3,null,3,null,1]
	Output: 7
	Explanation: Maximum amount of money the thief can rob = 3 + 3 + 1 = 7.

Example 2:

	Input: root = [3,4,5,1,3,null,1]
	Output: 9
	Explanation: Maximum amount of money the thief can rob = 4 + 5 = 9.


Constraints:

	The number of nodes in the tree is in the range [1, 10^4].
	0 <= Node.val <= 10^4
*/

// TreeNode definition
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rob(root *TreeNode) int {
	robRoot, skipRoot := dfs(root)
	return max(robRoot, skipRoot)
}

// dfs function returns two values: max money if robbing or not robbing current node
func dfs(node *TreeNode) (int, int) {
	if node == nil {
		return 0, 0
	}

	// Recursively solve the subproblems
	leftRob, leftSkip := dfs(node.Left)
	rightRob, rightSkip := dfs(node.Right)

	// If we rob this node, we cannot rob its children
	robThis := node.Val + leftSkip + rightSkip

	// If we don't rob this node, we can choose to rob or not rob its children
	skipThis := max(leftRob, leftSkip) + max(rightRob, rightSkip)

	return robThis, skipThis
}

// Utility function to get max of two integers
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	// Example 1
	root1 := &TreeNode{Val: 3, Left: &TreeNode{Val: 2, Right: &TreeNode{Val: 3}}, Right: &TreeNode{Val: 3, Right: &TreeNode{Val: 1}}}
	fmt.Println(rob(root1)) // Output: 7

	// Example 2
	root2 := &TreeNode{Val: 3, Left: &TreeNode{Val: 4, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 3}}, Right: &TreeNode{Val: 5, Right: &TreeNode{Val: 1}}}
	fmt.Println(rob(root2)) // Output: 9
}
