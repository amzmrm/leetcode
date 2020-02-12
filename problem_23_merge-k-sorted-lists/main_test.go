package problem_23_merge_k_sorted_lists

import (
	"math/rand"
	"reflect"
	"sort"
	"testing"
)

func TestMergeKLists(t *testing.T) {
	type test struct {
		name   string
		lists  []*ListNode
		expect []int
	}

	tests := make([]*test, 0)

	list1 := &ListNode{Val: 1}
	list1Node2 := &ListNode{Val: 4}
	list1Node3 := &ListNode{Val: 5}
	list1.Next = list1Node2
	list1Node2.Next = list1Node3

	list2 := &ListNode{Val: 1}
	list2Node2 := &ListNode{Val: 3}
	list2Node3 := &ListNode{Val: 4}
	list2.Next = list2Node2
	list2Node2.Next = list2Node3

	list3 := &ListNode{Val: 2}
	list3Node2 := &ListNode{Val: 6}
	list3.Next = list3Node2

	normals := make([]*ListNode, 0)
	normals = append(normals, list1, list2, list3)
	//tests = append(tests, &test{
	//	name:   "normal",
	//	lists:  normals,
	//	expect: []int{1, 1, 2, 3, 4, 4, 5, 6},
	//})

	singleList := make([]*ListNode, 0)
	singleList = append(singleList, list1)
	tests = append(tests, &test{
		name:   "single list",
		lists:  singleList,
		expect: []int{1, 4, 5},
	})

	emptyList := make([]*ListNode, 0)
	tests = append(tests, &test{
		name:   "empty list",
		lists:  emptyList,
		expect: []int{},
	})

	cornerCase1 := make([]*ListNode, 0)
	cornerCase1 = append(cornerCase1, list1)
	cornerCase1 = append(cornerCase1, nil)
	tests = append(tests, &test{
		name:   "corner case 1",
		lists:  cornerCase1,
		expect: []int{1, 4, 5},
	})

	N := 10000
	ints := make([]int, 0, N)
	for i := 0; i < N; i++ {
		ints = append(ints, int(rand.Int31n(int32(N))))
	}
	expect := make([]int, 0, N)
	cornerCase2 := make([]*ListNode, 0, N)
	for _, num := range ints {
		expect = append(expect, num)
		cornerCase2 = append(cornerCase2, &ListNode{Val: num})
	}

	sort.Ints(expect)

	tests = append(tests, &test{
		name:   "corner case 2",
		lists:  cornerCase2,
		expect: expect,
	})

	list4 := &ListNode{Val: 4}
	list5 := &ListNode{Val: 5}
	cornerCase3 := make([]*ListNode, 0)
	cornerCase3 = append(cornerCase3, list4)
	cornerCase3 = append(cornerCase3, list5)
	tests = append(tests, &test{
		name:   "corner case 3",
		lists:  cornerCase3,
		expect: []int{4, 5},
	})

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := MergeKLists(test.lists)
			if !reflect.DeepEqual(got.ToSlice(), test.expect) {
				t.Fatalf("got:%v, expect:%v", got, test.expect)
			}
		})
	}
}
