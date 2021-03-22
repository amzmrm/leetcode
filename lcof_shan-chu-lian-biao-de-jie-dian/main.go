package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteNode(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}
	if head.Val == val {
		return head.Next
	}
	pre, curr := head, head.Next
	for curr != nil && curr.Val != val {
		pre = curr
		curr = curr.Next
	}
	if curr != nil {
		pre.Next = curr.Next
	}
	return head
}
