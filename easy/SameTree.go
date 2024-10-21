// условие задачи - https://leetcode.com/problems/same-tree/description/

// time complexity - O(n)
// space complexity - O(n)
package main

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

// BFS
func isSameTreeBFS(p *TreeNode, q *TreeNode) bool {
    queue1 := []*TreeNode{p}
    queue2 := []*TreeNode{q}

    for len(queue1) > 0 && len(queue2) > 0 {
		if queue1[0] == nil && queue2[0] == nil {
			queue1 = queue1[1:]
			queue2 = queue2[1:]
			continue
		} else if queue1[0] == nil || queue2[0] == nil {
			return false
		}
        if queue1[0].Val != queue2[0].Val {
            return false
        }
		queue1 = append(queue1, queue1[0].Left, queue1[0].Right)
		queue2 = append(queue2, queue2[0].Left, queue2[0].Right)

		queue1 = queue1[1:]
		queue2 = queue2[1:]
    }
	
	if len(queue1) > 0 || len(queue2) > 0 {
		return false
	}
    return true
}

//DFS
func isSameTreeDFS(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
	  return true
	}
	if p == nil || q == nil || p.Val != q.Val {
	  return false
	}
  
	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}