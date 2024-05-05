// условие задачи - https://leetcode.com/problems/delete-node-in-a-linked-list/description

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

 func deleteNode(node *ListNode) {
    *node = *node.Next
}