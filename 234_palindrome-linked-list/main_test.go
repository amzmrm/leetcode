package leetcode

import (
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	type test struct {
		name   string
		list   *ListNode
		expect bool
	}

	tests := make([]test, 0)
	l1n1 := &ListNode{Val: 1}
	l1n2 := &ListNode{Val: 2}
	l1n3 := &ListNode{Val: 3}
	l1n3_2 := &ListNode{Val: 4}
	l1n4 := &ListNode{Val: 3}
	l1n5 := &ListNode{Val: 2}
	l1n6 := &ListNode{Val: 1}
	l1n1.Next = l1n2
	l1n2.Next = l1n3
	l1n3.Next = l1n4
	l1n3_2.Next = l1n4
	l1n4.Next = l1n5
	l1n5.Next = l1n6
	tests = append(tests, test{
		name:   "normal example 1",
		list:   l1n1,
		expect: true,
	})

	l2n1 := &ListNode{Val: 1}
	l2n2 := &ListNode{Val: 1}
	l2n1.Next = l2n2
	tests = append(tests, test{
		name:   "only-two-elements-list",
		list:   l2n1,
		expect: true,
	})

	l3n1 := &ListNode{Val: 1}
	tests = append(tests, test{
		name:   "only-one-elements-list",
		list:   l3n1,
		expect: true,
	})

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := IsPalindrome(test.list)
			if got != test.expect {
				t.Fatalf("test failed")
			}
		})
	}
}
