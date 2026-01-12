// условие задача - https://leetcode.com/problems/linked-list-cycle/description
// метод slow и fast
package main

type ListNode struct {
    Val int
    Next *ListNode
}

// метод slow и fast
func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	
	slow := head.Next
	fast := head.Next.Next

	for fast != nil || slow != nil || fast != slow {
		fast = head.Next.Next
		slow = slow.Next
	}

	if slow == fast {
		return true
	}

	return false
}

// мой очень медленный метод - O(n*n)
/*
func hasCycle(head *ListNode) bool {
	curr, prev := head, head
	
	for curr.Next != nil && curr.Next != prev {
		if curr == prev {
			curr = curr.Next
			prev = head
		} else {
			prev = prev.Next
		}
	}

	if curr.Next == nil {
		return false
	}

	return true
}
*/
