package problem_141_linked_list_cycle

import (
	"testing"
)

func TestHasCycle(t *testing.T) {
	type test struct {
		name   string
		list   *ListNode
		expect bool
	}

	normalCase := &ListNode{Val: 3}
	normalCaseNode2 := &ListNode{Val: 2}
	normalCaseNode3 := &ListNode{Val: 0}
	normaCaseNode4 := &ListNode{Val: 4}
	normalCase.Next = normalCaseNode2
	normalCaseNode2.Next = normalCaseNode3
	normalCaseNode3.Next = normaCaseNode4
	normaCaseNode4.Next = normalCaseNode2

	anotherCase := &ListNode{Val: 1}
	anotherCaseNode2 := &ListNode{Val: 2}
	anotherCase.Next = anotherCaseNode2
	anotherCaseNode2.Next = anotherCase

	singleNode := &ListNode{Val: 1}

	theThirdCase := &ListNode{Val:1}
	theThirdCase.Next = &ListNode{Val:2}

	var emptyList *ListNode

	var tests = []test{
		{
			name:   "normal case",
			list:   normalCase,
			expect: true,
		},
		{
			name:   "another case",
			list:   anotherCase,
			expect: true,
		},
		{
			name:   "single node",
			list:   singleNode,
			expect: false,
		},
		{
			name:   "empty list",
			list:   emptyList,
			expect: false,
		},
		{
			name:   "the third case",
			list:   theThirdCase,
			expect: false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := HasCycle(test.list)
			if got != test.expect {
				t.Fatalf("got:%v, expect:%v", got, test.expect)
			}
		})
	}
}
