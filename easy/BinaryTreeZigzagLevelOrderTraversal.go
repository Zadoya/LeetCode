// условие задачи - https://leetcode.com/problems/binary-tree-zigzag-level-order-traversal/?envType=problem-list-v2&envId=breadth-first-search

// time complexity - O(n)
// space complexity - O(n)
package main
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int { 
	if root == nil {
		return nil
	}
	var result [][]int
    queue := []*TreeNode{root}
	leftToRight := true

    for len(queue) > 0 {
		lvlSize := len(queue)
		lvl := make([]int, lvlSize)
		for i := 0; i < lvlSize; i++ {
			node := queue[0]
			queue = queue[1:] 

			if leftToRight {
				lvl[i] = node.Val
			} else {
				lvl[lvlSize-1-i] = node.Val
			}
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		result = append(result, lvl)
		leftToRight = !leftToRight
	}

	return result
}
