package leetcode

import (
	"reflect"
	"testing"
)

func TestReverseList(t *testing.T) {
	type test struct {
		list   *ListNode
		name   string
		expect []int
	}

	tests := make([]test, 0)

	// 正常
	list1 := &ListNode{Val: 1}
	l1Node2 := &ListNode{Val: 2}
	l1Node3 := &ListNode{Val: 3}
	l1Node4 := &ListNode{Val: 4}
	l1Node5 := &ListNode{Val: 5}
	list1.Next = l1Node2
	l1Node2.Next = l1Node3
	l1Node3.Next = l1Node4
	l1Node4.Next = l1Node5
	test1 := test{
		list:   list1,
		name:   "normal",
		expect: []int{5, 4, 3, 2, 1},
	}
	tests = append(tests, test1)

	// 空链表
	var list2 *ListNode
	tests = append(tests, test{
		list:   list2,
		name:   "empty",
		expect: []int{},
	})

	// 单节点链表
	list3 := &ListNode{Val: 1}
	tests = append(tests, test{
		list:   list3,
		name:   "length is 1",
		expect: []int{1},
	})

	// 恰好两个节点的链表
	list4 := &ListNode{Val: 1}
	l4Node2 := &ListNode{Val: 2}
	list4.Next = l4Node2
	tests = append(tests, test{
		list:   list4,
		name:   "length is 2",
		expect: []int{2, 1},
	})

	for _, test := range tests {
		got := ReverseList2(test.list)
		if !reflect.DeepEqual(got.ToSlice(), test.expect) {
			t.Fatalf("got:%#v, expect:%#v", got.ToSlice(), test.expect)
		}
	}
}
