package leetcode

import (
	"reflect"
	"testing"
)

func TestRemoveNthFromEnd(t *testing.T) {
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

		expect := []int{1, 3, 4, 5}

		got := RemoveNthFromEnd(&head, 4)
		if !reflect.DeepEqual(got.ToSlice(), expect) {
			t.Fatalf("got:%#v, expect:%#v", got.ToSlice(), expect)
		}
	})

	t.Run("nil", func(t *testing.T) {
		var head *ListNode
		got := RemoveNthFromEnd(head, 4)
		if got != nil {
			t.Fatalf("got:%#v, expect:%#v", got.ToSlice(), nil)
		}
	})

	t.Run("remove the last node", func(t *testing.T) {
		head := ListNode{Val: 1}
		node1 := ListNode{Val: 2}
		node2 := ListNode{Val: 3}
		node3 := ListNode{Val: 4}
		node4 := ListNode{Val: 5}
		head.Next = &node1
		node1.Next = &node2
		node2.Next = &node3
		node3.Next = &node4

		expect := []int{1, 2, 3, 4}

		got := RemoveNthFromEnd(&head, 1)
		if !reflect.DeepEqual(got.ToSlice(), expect) {
			t.Fatalf("got:%#v, expect:%#v", got.ToSlice(), expect)
		}
	})

	t.Run("remove the first node", func(t *testing.T) {
		head := ListNode{Val: 1}
		node1 := ListNode{Val: 2}
		node2 := ListNode{Val: 3}
		node3 := ListNode{Val: 4}
		node4 := ListNode{Val: 5}
		head.Next = &node1
		node1.Next = &node2
		node2.Next = &node3
		node3.Next = &node4

		expect := []int{2, 3, 4, 5}

		got := RemoveNthFromEnd(&head, 5)
		if !reflect.DeepEqual(got.ToSlice(), expect) {
			t.Fatalf("got:%#v, expect:%#v", got.ToSlice(), expect)
		}
	})
}
