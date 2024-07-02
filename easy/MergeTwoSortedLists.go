// условие задачи - https://leetcode.com/problems/merge-two-sorted-lists/description/
type ListNode struct {
    Val int
    Next *ListNode
}

 func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    head := &ListNode{} 

    curr := head
    for list1 != nil && list2 != nil {
        if list1.Val <= list2.Val {
            curr.Next = list1 
            curr = curr.Next
            list1 = list1.Next
        } else {
            curr.Next = list2
            curr = curr.Next
            list2 = list2.Next
        }
    }

    if list1 != nil {
		curr.Next = list1
    }
    if list2 != nil {
		curr.Next = list2
    }

    return head.Next
}