// условие задачи - https://leetcode.com/problems/convert-sorted-array-to-binary-search-tree/description/

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	var root *TreeNode
   
	left := 0
	right := len(nums) - 1
	if left <= right {
		mid := (left + right + 1) / 2
		root = &TreeNode{
			Val: nums[mid],
			Left: sortedArrayToBST(nums[:mid]),
			Right: sortedArrayToBST(nums[mid + 1:]),
		}
	}
	return root
}