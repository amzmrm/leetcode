package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	sentinel := &ListNode{Next: head}
	pre, curr := sentinel, head
	for curr != nil {
		if curr.Val == val {
			pre.Next = curr.Next
		} else {
			pre = curr
			curr = curr.Next
		}
	}
	return sentinel.Next
}
