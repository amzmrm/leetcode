package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func RotateList(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}

	if k == 0 {
		return head
	}

	var length int
	for p := head; p != nil; p = p.Next {
		length++
	}

	first, second := head, head

	steps := k % length
	for i := 0; i < steps; i++ {
		second = second.Next
	}

	for second.Next != nil {
		first = first.Next
		second = second.Next
	}

	second.Next = head
	newHead := first.Next
	first.Next = nil

	return newHead
}
