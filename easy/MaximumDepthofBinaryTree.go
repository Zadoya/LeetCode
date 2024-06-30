// условие задачи - https://leetcode.com/problems/maximum-depth-of-binary-tree/description

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
    if root != nil {
       return 1 + max(maxDepth(root.Left), maxDepth(root.Right)) 
    }
    return 0   
}