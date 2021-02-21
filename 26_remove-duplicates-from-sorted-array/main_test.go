package leetcode

import (
	"reflect"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	type test struct {
		name   string
		nums   []int
		expect []int
	}

	tests := make([]test, 0)
	tests = append(tests, test{
		name:   "leetcode example 1",
		nums:   []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
		expect: []int{0, 1, 2, 3, 4},
	})
	tests = append(tests, test{
		name:   "leetcode example 2",
		nums:   []int{1, 1, 2},
		expect: []int{1, 2},
	})
	tests = append(tests, test{
		name:   "single element",
		nums:   []int{1},
		expect: []int{1},
	})
	tests = append(tests, test{
		name:   "empty array",
		nums:   []int{},
		expect: []int{},
	})
	tests = append(tests, test{
		name:   "all elements art distinct",
		nums:   []int{1, 2, 3, 4, 5, 6},
		expect: []int{1, 2, 3, 4, 5, 6},
	})

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			length := RemoveDuplicates2(test.nums)
			if !reflect.DeepEqual(test.nums[:length], test.expect) {
				t.Fatalf("got:%v, expect:%v", test.nums, test.expect)
			}
		})
	}
}
