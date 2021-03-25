package leetcode

import "testing"

func TestName(t *testing.T) {
	head := &ListNode{Val: 1, Next: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3}}}}
	head = deleteDuplicates(head)
}
