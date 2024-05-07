//условие задачи - https://leetcode.com/problems/reverse-linked-list/description/
/**
* type ListNode struct {
*    Val int
*    Next *ListNode
* }
*/

func reverseList(head *ListNode) *ListNode {
	var reversed *ListNode

	for head != nil {
		reversed = &ListNode{Val: head.Val, Next: reversed}
		head = head.Next
	}
	return reversed
}