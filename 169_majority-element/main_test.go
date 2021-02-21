package leetcode

import "testing"

func TestMajorityElement(t *testing.T) {
	type test struct {
		name   string
		nums   []int
		expect int
	}

	tests := make([]*test, 0)
	tests = append(tests, &test{
		name:   "normal case 1",
		nums:   []int{2, 2, 1, 1, 1, 2, 2},
		expect: 2,
	})
	tests = append(tests, &test{
		name:   "normal case 2",
		nums:   []int{3, 2, 3},
		expect: 3,
	})

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := majorityElement(test.nums)
			if got != test.expect {
				t.Fatalf("got:%v, expect:%v", got, test.expect)
			}
		})
	}
}
