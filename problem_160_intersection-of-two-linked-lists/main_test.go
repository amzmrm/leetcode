package main

import "testing"

func TestGetIntersectionNode(t *testing.T) {
	t.Run("normal", func(t *testing.T) {
		interRoot := ListNode{Val: 8}
		interNode1 := ListNode{Val: 4}
		interNode2 := ListNode{Val: 5}
		interRoot.Next = &interNode1
		interNode1.Next = &interNode2

		aRoot := ListNode{Val: 4}
		aNode1 := ListNode{Val: 1}
		aRoot.Next = &aNode1
		aNode1.Next = &interRoot

		bRoot := ListNode{Val: 5}
		bNode1 := ListNode{Val: 0}
		bNode2 := ListNode{Val: 1}
		bRoot.Next = &bNode1
		bNode1.Next = &bNode2
		bNode2.Next = &interRoot

		expect := ListNode{Val: 8}

		got := GetIntersectionNode(&aRoot, &bRoot)
		if got == nil {
			t.Fatalf("expect:%v, got:nil", expect.Val)
		}
		if got.Val != expect.Val {
			t.Fatalf("expect:%v, got:%v", expect.Val, got.Val)
		}
	})

	t.Run("no intersection", func(t *testing.T) {
		aRoot := ListNode{Val: 2}
		aNode1 := ListNode{Val: 6}
		aNode2 := ListNode{Val: 4}
		aRoot.Next = &aNode1
		aNode1.Next = &aNode2

		bRoot := ListNode{Val: 1}
		bNode1 := ListNode{Val: 5}
		bRoot.Next = &bNode1

		var expect *ListNode

		got := GetIntersectionNode(&aRoot, &bRoot)
		if got != expect {
			t.Fatalf("expect:nil, got:%v", got.Val)
		}
	})
}
