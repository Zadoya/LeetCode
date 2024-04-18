// Условие задачи - https://leetcode.com/problems/add-one-row-to-tree/description/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func addNewRoot(root *TreeNode, val int) *TreeNode {
	return &TreeNode{Val: val, Left: root}
}

func addOneRow(root *TreeNode, val int, depth int) *TreeNode {
	if depth == 1 {
		return addNewRoot(root, val)
	}
	if depth - 1 == 1 {
		newLeft := &TreeNode{Val: val, Left: root.Left}
        newRight := &TreeNode{Val: val, Right: root.Right}
        root.Left = newLeft
        root.Right = newRight
	} 
    if root.Left != nil {
		addOneRow(root.Left, val, depth - 1)
	}
	if root.Right != nil {
		addOneRow(root.Right, val, depth - 1)
	}
    return root
}