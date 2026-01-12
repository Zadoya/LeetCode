// условие задачи - https://leetcode.com/problems/minimum-absolute-difference-in-bst/description

package main

type TreeNode struct {
    Val int
    Left *TreeNode
	Right *TreeNode
}

func LNR(root *TreeNode, res *[]int) {
	if root != nil {
		LNR(root.Left, res)
		*res = append(*res, root.Val)
		LNR(root.Right, res)
	}
}

func getMinimumDifference(root *TreeNode) int {
	res := make([]int, 0)
	LNR(root, &res)
	
	if len(res) < 2 {
		return 0
	}
	
	minDiff := res[1] - res[0]
	for i := 2; i < len(res); i++ {
		minDiff = min(minDiff, res[i] - res[i - 1])
	}
	return minDiff
}
