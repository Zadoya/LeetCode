// условие задачи - https://leetcode.com/problems/double-a-number-represented-as-a-linked-list/descriptio

/*
type ListNode struct {
    Val int
    Next *ListNode
}
*/

func doubleIt(head *ListNode) *ListNode {
	curr := head
	
    if curr.Val > 4 {
        head = &ListNode{Val: 1, Next: head}
    }

	for curr != nil {
		curr.Val = (curr.Val * 2) % 10
        if curr.Next != nil && curr.Next.Val > 4 {
            curr.Val += 1
        }
		curr = curr.Next
	}

	return head
}