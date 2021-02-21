package leetcode

import "testing"

func TestFib(t *testing.T) {
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
		name:   "5",
		N:      5,
		expect: 5,
	})
	tests = append(tests, test{
		name:   "8",
		N:      8,
		expect: 21,
	})

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := Fib(test.N)
			if got != test.expect {
				t.Fatalf("got:%v, want:%v", got, test.expect)
			}
		})
	}

}
