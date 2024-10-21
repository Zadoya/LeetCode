//условие задачи - https://leetcode.com/problems/serialize-and-deserialize-bst/description/

package main

import (
	"fmt"
	"strconv"
	"strings"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Codec struct {

}

func Constructor() Codec {
    return Codec{}
}
// DFS метод, быстрее чем BFS, но по временной сложности и по памяти одинаковые
// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return ""
	}
	return strconv.FormatUint(uint64(root.Val), 10) + "|" + this.serialize(root.Left) + this.serialize(root.Right)
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {  
	strs := strings.Split(data, "|")
	return this.helper(&strs, math.MinInt64, math.MaxInt64)
}

func (this *Codec) helper(data *[]string, left, right int64) *TreeNode {
	if len(*data) == 0 || (*data)[0] == "" {
		return nil
	}
	num, err := strconv.ParseInt((*data)[0], 10, 64)
	if num < left || num > right || err != nil {
		return nil
	}
	*data = (*data)[1:]
	root := &TreeNode{Val: int(num)}
	root.Left = this.helper(data, left, num)
	root.Right = this.helper(data, num, right)
	return root
}

// BFS метод
/*
func (this *Codec) serialize(root *TreeNode) string {
	result := ""
	if root == nil {
		return result
	}
	queue := []*TreeNode{root}
	
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if node == nil {
			result += "|"
		} else {
			result += strconv.FormatUint(uint64(node.Val), 10) + "|"
            queue = append(queue, node.Left, node.Right)
		}
	}
	return result[:len(result) - 1]
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {  
    if len(data) == 0 {
        return nil
    }  
	strs := strings.Split(data, "|")
	nums := make([]int, 0, len(strs))
	for _, numStr := range strs {
		if len(numStr) == 0{
			nums = append(nums, -1)
		} else {
			num, _ := strconv.ParseUint(numStr, 10, 32)
			nums = append(nums, int(num))
		}
	}
	root := &TreeNode{Val: nums[0]}
	fmt.Println(nums)
	queue := []*TreeNode{root}
	leftOrRight := true

	for i := 1; len(queue) > 0; i++ {
		node := queue[0]
		queue = queue[1:]
		addNum := func (node *TreeNode, nums []int) {
			if nums[i] >= 0 {
				if leftOrRight {
					node.Left = &TreeNode{Val: nums[i]}
					queue = append(queue, node.Left)
					leftOrRight = !leftOrRight
				} else {
					node.Right = &TreeNode{Val: nums[i]}
					queue = append(queue, node.Right)
					leftOrRight = !leftOrRight
				}
			} else {
				leftOrRight = !leftOrRight
			}
		}
		addNum(node, nums)
		i++
		addNum(node, nums)
	}
	return root
}
*/