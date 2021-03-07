package leetcode

import "testing"

func TestPermute(t *testing.T) {
	got := permute([]int{1, 2, 3})
	if len(got) != 6 {
		t.Fatalf("expect: %d, got: %d", 6, len(got))
	}
}
