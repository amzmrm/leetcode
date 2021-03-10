package leetcode

import "testing"

func TestPlusOne(t *testing.T) {
	index := 0
	tests := []struct {
		input  string
		expect string
	}{
		{
			input:  "0000",
			expect: "1000",
		},
		{
			input:  "9000",
			expect: "0000",
		},
		{
			input:  "5000",
			expect: "6000",
		},
	}

	for _, test := range tests {
		got := plusOne(test.input, index)
		if got != test.expect {
			t.Fatalf("expect: %s, got: %s", test.expect, got)
		}
	}
}

func BenchmarkPlusOne(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := "0000"
		index := 0
		got := plusOne(s, index)
		if got != "9000" {
			b.Fatalf("expect: %s, got: %s", "9000", got)
		}
	}
}

func TestMinusOne(t *testing.T) {
	index := 0
	tests := []struct {
		input  string
		expect string
	}{
		{
			input:  "0000",
			expect: "9000",
		},
		{
			input:  "5000",
			expect: "4000",
		},
		{
			input:  "1000",
			expect: "0000",
		},
	}

	for _, test := range tests {
		got := minusOne(test.input, index)
		if got != test.expect {
			t.Fatalf("expect: %s, got: %s", test.expect, got)
		}
	}
}
