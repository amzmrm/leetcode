package leetcode

import (
	"log"
	"testing"
)

func TestMergeTwoLists(t *testing.T) {
	n1 := &ListNode{Val: 1}
	n2 := &ListNode{Val: 2}
	n3 := &ListNode{Val: 4}
	n31 := &ListNode{Val: 7}
	n32 := &ListNode{Val: 8}
	n1.Next = n2
	n2.Next = n3
	n3.Next = n31
	n31.Next = n32

	n4 := &ListNode{Val: 1}
	n5 := &ListNode{Val: 3}
	n6 := &ListNode{Val: 4}
	n7 := &ListNode{Val: 5}
	n4.Next = n5
	n5.Next = n6
	n6.Next = n7

	l1, l2 := n1, n4

	merged := MergeTwoLists(l1, l2)
	for node := merged; node != nil; {
		log.Println(node.Val)
		node = node.Next
	}
}
