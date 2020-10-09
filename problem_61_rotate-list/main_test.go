package main

import (
	"testing"
)

func TestRotateList(t *testing.T) {
	t.Run("normal", func(t *testing.T) {
		head := ListNode{Val: 1}
		node1 := ListNode{Val: 2}
		node2 := ListNode{Val: 3}
		node3 := ListNode{Val: 4}
		node4 := ListNode{Val: 5}
		head.Next = &node1
		node1.Next = &node2
		node2.Next = &node3
		node3.Next = &node4

		got := RotateList(&head, 2)
		if got == nil {
			t.Fatalf("expect:%v, got:%v", 4, nil)
		}

		expect := node3
		if got.Val != expect.Val {
			t.Fatalf("expect:%v, got:%v", expect.Val, got.Val)
		}
	})

	t.Run("k greater than list length", func(t *testing.T) {
		head := ListNode{Val: 0}
		node1 := ListNode{Val: 1}
		node2 := ListNode{Val: 2}
		head.Next = &node1
		node1.Next = &node2

		got := RotateList(&head, 4)
		if got == nil {
			t.Fatalf("expect:%v, got:%v", node2.Val, nil)
		}

		expect := node2
		if got.Val != expect.Val {
			t.Fatalf("expect:%v, got:%v", expect.Val, got.Val)
		}
	})

	t.Run("single elem", func(t *testing.T) {
		head := ListNode{Val: 1}
		got := RotateList(&head, 0)
		expect := head
		if got.Val != expect.Val {
			t.Fatalf("expect:%v, got:%v", expect.Val, got.Val)
		}
	})

	t.Run("rotate all", func(t *testing.T) {
		head := ListNode{Val: 1}
		got := RotateList(&head, 1)
		expect := head
		if got.Val != expect.Val {
			t.Fatalf("expect:%v, got:%v", expect.Val, got.Val)
		}
	})

	t.Run("another rotate all", func(t *testing.T) {
		head := ListNode{Val: 1}
		node1 := ListNode{Val: 2}
		node2 := ListNode{Val: 3}
		node3 := ListNode{Val: 4}
		node4 := ListNode{Val: 5}
		head.Next = &node1
		node1.Next = &node2
		node2.Next = &node3
		node3.Next = &node4

		got := RotateList(&head, 5)
		expect := head
		if got.Val != expect.Val {
			t.Fatalf("expect:%v, got:%v", expect.Val, got.Val)
		}
	})
}
