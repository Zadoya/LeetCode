// условие задачи - https://leetcode.com/problems/binary-tree-inorder-traversal/

package main

type TreeNode struct {
	Val int
	Left *TreeNode
 	Right *TreeNode
}
 

/* рекурсивный метод

func inorderTraversal(root *TreeNode) []int {
    var res []int

    if root != nil {
        res = append(res, inorderTraversal(root.Left)...)
        res = append(res, root.Val)
        res = append(res, inorderTraversal(root.Right)...)
        
    } 
    return res   
}
*/

// итеративный со стеком
/*
func inorderTraversal(root *TreeNode) []int {
    var (
        res []int
        stack []*TreeNode
    )
    curr := root

    for curr != nil || len(stack) != 0 {
        for curr != nil{
            stack = append(stack, curr)
            curr = curr.Left
        }
        curr = stack[len(stack) - 1]
        stack = stack[:len(stack) - 1]
        res = append(res, curr.Val)
        curr = curr.Right
    } 
    return res   
}
*/

// обход Морриса - Morris Traversal

func inorderTraversal(root *TreeNode) []int {
    var res []int

    curr := root

    for curr != nil {
        if curr.Left == nil {
            res = append(res, curr.Val)
            curr = curr.Right
        } else {
            subtree := curr.Left
            for subtree.Right != nil {
                subtree = subtree.Right
            }
            subtree.Right = curr
            curr = curr.Left
            subtree.Right.Left = nil
        }
    }

    return res
}