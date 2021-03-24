package leetcode

import "testing"

func TestName(t *testing.T) {
	root := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}}
	reorderList(root)
	t.Log(root)
}
