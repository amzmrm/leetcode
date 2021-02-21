package leetcode

import "testing"

func TestClimbStairs(t *testing.T) {
	type test struct {
		name   string
		N      int
		expect int
	}

	tests := make([]test, 0)
	tests = append(tests, test{
		name:   "1",
		N:      1,
		expect: 1,
	})
	tests = append(tests, test{
		name:   "2",
		N:      2,
		expect: 2,
	})
	tests = append(tests, test{
		name:   "3",
		N:      3,
		expect: 3,
	})
	tests = append(tests, test{
		name:   "4",
		N:      4,
		expect: 5,
	})
	tests = append(tests, test{
		name:   "5",
		N:      5,
		expect: 8,
	})

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := ClimbStairs(test.N)
			if got != test.expect {
				t.Fatalf("got:%v, want:%v", got, test.expect)
			}
		})
	}
}
