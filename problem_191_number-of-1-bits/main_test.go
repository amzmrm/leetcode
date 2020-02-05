package problem_191_number_of_1_bits

import "testing"

func TestHammingWeight(t *testing.T) {
	type test struct {
		name   string
		num    uint32
		expect int
	}

	tests := make([]test, 0)
	tests = append(tests, test{
		name:   "2",
		num:    2,
		expect: 1,
	})
	tests = append(tests, test{
		name:   "3",
		num:    3,
		expect: 2,
	})
	tests = append(tests, test{
		name:   "99999999",
		num:    99999999,
		expect: 19,
	})

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := HammingWeight2(test.num)
			if got != test.expect {
				t.Fatalf("got:%v, expect:%v", got, test.expect)
			}
		})
	}
}
