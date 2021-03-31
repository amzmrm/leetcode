package leetcode

import "testing"

func TestSortedListToBST(t *testing.T) {
	test := &ListNode{Val: -10, Next: &ListNode{Val: 3, Next: &ListNode{Val: 0, Next: &ListNode{Val: 5, Next: &ListNode{Val: 9}}}}}
	got := sortedListToBST(test)
	t.Log(got.Val)
}
