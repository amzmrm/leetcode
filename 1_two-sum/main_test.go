package leetcode

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	type test struct {
		name   string
		list   []int
		target int
		expect []int
	}

	tests := make([]test, 0)
	tests = append(tests, test{
		name:   "normal",
		list:   []int{2, 7, 11, 15},
		target: 9,
		expect: []int{0, 1},
	})

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := TwoSum(test.list, test.target)
			if !reflect.DeepEqual(got, test.expect) {
				t.Fatalf("got:%#v, expect:%#v", got, test.expect)
			}
		})
	}
}
