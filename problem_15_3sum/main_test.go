package main

import (
	"reflect"
	"testing"
)

func TestThreeSum(t *testing.T) {
	type test struct {
		name   string
		nums   []int
		expect [][]int
	}

	tests := make([]*test, 0)

	list1 := []int{-1, 0, 1, 2, -1, -4}
	tests = append(tests, &test{
		name:   "normal case 1",
		nums:   list1,
		expect: [][]int{{-1, -1, 2}, {-1, 0, 1}},
	})

	list2 := []int{0, 0, 0}
	tests = append(tests, &test{
		name:   "all zeros",
		nums:   list2,
		expect: [][]int{{0, 0, 0}},
	})

	list3 := []int{1,2,3,4,5,6}
	tests = append(tests, &test{
		name:   "starting with 1",
		nums:   list3,
		expect: [][]int{},
	})

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := ThreeSum(test.nums)
			if !reflect.DeepEqual(got, test.expect) {
				t.Fatalf("got:%v, expect:%v", got, test.expect)
			}
		})
	}
}
