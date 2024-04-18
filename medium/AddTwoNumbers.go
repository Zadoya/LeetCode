// Условие задачи - https://leetcode.com/problems/add-two-numbers/description/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	result := &ListNode{}
	current := result
	dec := 0
	
	for l1 != nil || l2 != nil || dec != 0 {
		if l1 != nil {
            dec += l1.Val
            l1 = l1.Next
        }
        if l2 != nil {
			dec += l2.Val
			l2 = l2.Next
		}
        current.Val = dec % 10
		dec = dec / 10
		if l1 != nil || l2 != nil || dec != 0 {
			current.Next = &ListNode{}
			current = current.Next
		}
	}

	return result
}
