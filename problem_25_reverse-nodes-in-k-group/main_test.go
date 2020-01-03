package problem_25_reverse_nodes_in_k_group

import (
	"reflect"
	"testing"
)

func TestReverseKGroup(t *testing.T) {
	type test struct {
		name   string
		list   *ListNode
		k      int
		expect []int
	}

	tests := make([]test, 0)

	list1 := &ListNode{Val: 1}
	l1Node2 := &ListNode{Val: 2}
	l1Node3 := &ListNode{Val: 3}
	l1Node4 := &ListNode{Val: 4}

	list1.Next = l1Node2
	l1Node2.Next = l1Node3
	l1Node3.Next = l1Node4
	tests = append(tests, test{
		list:   list1,
		name:   "list length is a factor of k",
		k:      2,
		expect: []int{2, 1, 4, 3},
	})

	list2 := &ListNode{Val: 1}
	l2Node2 := &ListNode{Val: 2}
	l2Node3 := &ListNode{Val: 3}
	l2Node4 := &ListNode{Val: 4}
	l2Node5 := &ListNode{Val: 5}
	list2.Next = l2Node2
	l2Node2.Next = l2Node3
	l2Node3.Next = l2Node4
	l2Node4.Next = l2Node5
	tests = append(tests, test{
		name:   "list length is not a factor of k",
		list:   list2,
		k:      3,
		expect: []int{3, 2, 1, 4, 5},
	})

	var list3 *ListNode
	tests = append(tests, test{
		name:   "empty list",
		list:   list3,
		k:      2,
		expect: []int{},
	})

	list4 := &ListNode{Val: 1}
	l4Node2 := &ListNode{Val:2}
	list4.Next =l4Node2
	tests = append(tests, test{
		name:   "list length is smaller than k",
		list:   list4,
		k:      3,
		expect: []int{1, 2},
	})

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := ReverseKGroup(test.list, test.k).ToSlice()
			if !reflect.DeepEqual(got, test.expect) {
				t.Fatalf("got:%#v, expect:%#v", got, test.expect)
			}
		})
	}
}
