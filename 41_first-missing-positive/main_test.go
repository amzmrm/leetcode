package leetcode

import "testing"

func TestFirstMissingPositive(t *testing.T) {
	type test struct {
		name   string
		nums   []int
		expect int
	}

	tests := make([]*test, 0)

	list1 := []int{1, 2, 0}
	list2 := []int{1, 2, 3, 4, 5, 6}
	list3 := []int{3, 4, -1, 1}
	list4 := []int{7, 8, 9, 11, 12}

	tests = append(tests, &test{
		name:   "list1",
		nums:   list1,
		expect: 3,
	})

	tests = append(tests, &test{
		name:   "list2",
		nums:   list2,
		expect: 7,
	})

	tests = append(tests, &test{
		name:   "list3",
		nums:   list3,
		expect: 2,
	})

	tests = append(tests, &test{
		name:   "list4",
		nums:   list4,
		expect: 1,
	})

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := firstMissingPositive(test.nums)
			if got != test.expect {
				t.Fatalf("got:%v, expect:%v", got, test.expect)
			}
		})
	}
}
