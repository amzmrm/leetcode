package problem_2_add_two_numbers

import (
	"reflect"
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	type test struct {
		list1, list2 *ListNode
		name         string
		expect       []int
	}

	tests := make([]test, 0)

	// 正常输入，两个链表的长度一样
	list11 := &ListNode{Val: 2}
	l1Node12 := &ListNode{Val: 4}
	l1Node13 := &ListNode{Val: 3}
	list11.Next = l1Node12
	l1Node12.Next = l1Node13

	list12 := &ListNode{Val: 5}
	l2Node2 := &ListNode{Val: 6}
	l2Node3 := &ListNode{Val: 4}
	list12.Next = l2Node2
	l2Node2.Next = l2Node3

	tests = append(tests, test{
		list1:  list11,
		list2:  list12,
		name:   "normal",
		expect: []int{7, 0, 8},
	})

	// 两个链表的长度不一样
	list21 := &ListNode{Val: 2}
	node21 := &ListNode{Val: 4}
	node22 := &ListNode{Val: 3}
	node23 := &ListNode{Val: 9}
	list21.Next = node21
	node21.Next = node22
	node22.Next = node23

	list22 := &ListNode{Val: 5}
	node24 := &ListNode{Val: 6}
	list22.Next = node24

	tests = append(tests, test{
		list1:  list21,
		list2:  list22,
		name:   "length of list not equal to each other",
		expect: []int{7, 0, 4, 9},
	})

	// 两个链表的长度都是1，且连个节点的和值小于10
	list31 := &ListNode{Val: 5}
	list32 := &ListNode{Val: 4}
	tests = append(tests, test{
		list1:  list31,
		list2:  list32,
		name:   "both lists have only one element, and sum is not greater than 9",
		expect: []int{9},
	})

	// 两个链表的长度都是1，且两个节点的和值等于10
	list41 := &ListNode{Val: 5}
	list42 := &ListNode{Val: 5}
	tests = append(tests, test{
		list1:  list41,
		list2:  list42,
		name:   "both lists have only one element, and sum is 10",
		expect: []int{0, 1},
	})

	// 连个链表的长度都是1，且两个节点的和值大于10
	list51 := &ListNode{Val: 9}
	list52 := &ListNode{Val: 9}
	tests = append(tests, test{
		list1:  list51,
		list2:  list52,
		name:   "both lists have only one element, and sum is greater than 10",
		expect: []int{8, 1},
	})

	// 其中一个链表的长度为0
	list61 := &ListNode{Val: 8}
	list61.Next = &ListNode{Val: 9}
	var list62 *ListNode
	list62 = nil
	tests = append(tests, test{
		list1:  list61,
		list2:  list62,
		name:   "one of list is empty",
		expect: []int{8, 9},
	})

	for _, test := range tests {
		got := AddTwoNumbers(test.list1, test.list2)
		if !reflect.DeepEqual(got.ToSlice(), test.expect) {
			t.Fatalf("got: %#v, expect: %#v", got.ToSlice(), test.expect)
		}
	}
}
