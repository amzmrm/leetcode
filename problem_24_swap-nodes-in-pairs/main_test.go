package problem_24_swap_nodes_in_pairs

import (
	"reflect"
	"testing"
)

func TestSwapPairs(t *testing.T) {
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
	l1Node6 := &ListNode{Val: 6}
	list1.Next = l1Node2
	l1Node2.Next = l1Node3
	l1Node3.Next = l1Node4
	l1Node4.Next = l1Node5
	l1Node5.Next = l1Node6
	test1 := test{
		list:   list1,
		name:   "normal",
		expect: []int{2, 1, 4, 3, 6, 5},
	}
	tests = append(tests, test1)

	// length为单数
	list2 := &ListNode{Val: 1}
	l2Node2 := &ListNode{Val: 2}
	l2Node3 := &ListNode{Val: 3}
	l2Node4 := &ListNode{Val: 4}
	l2Node5 := &ListNode{Val: 5}
	list2.Next = l2Node2
	l2Node2.Next = l2Node3
	l2Node3.Next = l2Node4
	l2Node4.Next = l2Node5
	test2 := test{
		list:   list2,
		name:   "length is an odd number",
		expect: []int{2, 1, 4, 3, 5},
	}
	tests = append(tests, test2)

	// 单链表为空
	var list3 *ListNode
	test3 := test{
		list:   list3,
		name:   "empty list",
		expect: []int{},
	}
	tests = append(tests, test3)

	// 恰好两个节点的链表
	list4 := &ListNode{Val: 1}
	l4Node2 := &ListNode{Val: 2}
	list4.Next = l4Node2
	tests = append(tests, test{
		list:   list4,
		name:   "length is 2",
		expect: []int{2,1},
	})

	// 恰好1个节点的链表
	list5 := &ListNode{Val: 1}
	tests = append(tests, test{
		list:   list5,
		name:   "length is 1",
		expect: []int{1},
	})

	for _, test := range tests {
		got := SwapPairs(test.list)
		if !reflect.DeepEqual(got.ToSlice(), test.expect) {
			t.Fatalf("got:%#v, expect:%#v", got.ToSlice(), test.expect)
		}
	}
}
