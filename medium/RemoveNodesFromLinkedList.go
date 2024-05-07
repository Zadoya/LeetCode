// условие задачи - 

/*
type ListNode struct {
    Val int
    Next *ListNode
}
*/



/*** my solution *****

func reverseList(head *ListNode) *ListNode {
	var stack *ListNode

	for head != nil {
		stack = &ListNode{Val: head.Val, Next: reversed}
		head = head.Next
	}
	return stack
}

func removeNodes(head *ListNode) *ListNode {
	head = reverseList(head)
	max := 0
	var newList *ListNode

	for head != nil {
		if head.Val >= max {
			max = head.Val
			newList = &ListNode{Val: head.Val, Next: newList}
		}
		head = head.Next
	}

	return newList
}
*/


// ** find better solutuin on pointers ****
func reverseList(head *ListNode) *ListNode {
	var prev, curr *ListNode = nil, head

	for curr != nil {
		curr, prev, curr.Next = curr.Next, curr, prev 
	}
	return prev
}

func removeNodes(head *ListNode) *ListNode {
	head = reverseList(head)
	max := 0
	var prev, curr *ListNode = nil, head

	for curr != nil {
		if curr.Val >= max {
			max = curr.Val
			prev = curr
		} else {
			if prev != nil {
				prev.Next = curr.Next
			}
		}
		curr = curr.Next
	}

	return reverseList(head)
}