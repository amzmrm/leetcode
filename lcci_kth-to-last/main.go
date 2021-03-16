package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func kthToLast(head *ListNode, k int) int {
	if head == nil {
		return -1
	}
	fast, slow := head, head
	for fast != nil && k > 0 {
		fast = fast.Next
		k--
	}
	for fast != nil {
		slow = slow.Next
		fast = fast.Next
	}
	return slow.Val
}
