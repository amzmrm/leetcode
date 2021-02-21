package leetcode

import (
	"reflect"
	"testing"
)

func TestReversePrint(t *testing.T) {
	type test struct {
		name   string
		list   *ListNode
		expect []int
	}

	normal := &ListNode{Val: 1}
	normalNode2 := &ListNode{Val: 3}
	normalNode3 := &ListNode{Val: 2}
	normal.Next = normalNode2
	normalNode2.Next = normalNode3

	singleNode := &ListNode{Val: 1}

	tests := []test{
		{
			name:   "normal case",
			list:   normal,
			expect: []int{2, 3, 1},
		},
		{
			name:   "single node",
			list:   singleNode,
			expect: []int{1},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := ReversePrint(test.list)
			if !reflect.DeepEqual(got, test.expect) {
				t.Fatalf("got: %v, expect: %v", got, test.expect)
			}
		})
	}
}
