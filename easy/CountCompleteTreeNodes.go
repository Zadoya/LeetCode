//условие задачи - https://leetcode.com/problems/count-complete-tree-nodes/description/
package main

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func countNodes(root *TreeNode) int {
	count := 0
	if root == nil {
		return 0
	} else {
		count = 1
	}

	if root.Left != nil {
		count += countNodes(root.Left)
	}
	if root.Right != nil {
		count += countNodes(root.Right)
	}

	return count
}

// идеальное решение
/*
func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return 1 + countNodes(root.Left) + countNodes(root.Right
}
*/
