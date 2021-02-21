package leetcode

import "testing"

func TestFindMin(t *testing.T) {
	type test struct {
		name   string
		nums   []int
		expect int
	}

	tests := make([]*test, 0)
	tests = append(tests, &test{
		name:   "length is 1",
		nums:   []int{1},
		expect: 1,
	})
	tests = append(tests, &test{
		name:   "length is 2",
		nums:   []int{2, 1},
		expect: 1,
	})
	tests = append(tests, &test{
		name:   "0 nums rotated",
		nums:   []int{1, 2, 3, 4, 5},
		expect: 1,
	})
	tests = append(tests, &test{
		name:   "mid location is uncertain case 1",
		nums:   []int{2, 1, 2, 2, 2},
		expect: 1,
	})
	tests = append(tests, &test{
		name:   "mid location is uncertain case 2",
		nums:   []int{2, 1, 2, 2, 2},
		expect: 1,
	})
	tests = append(tests, &test{
		name:   "normal",
		nums:   []int{3, 4, 5, 1, 2},
		expect: 1,
	})

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := FindMin(test.nums)
			if got != test.expect {
				t.Fatalf("test %s failed, expect: %d, got: %d", test.name, test.expect, got)
			}
		})
	}
}
