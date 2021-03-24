package leetcode

import "testing"

func TestName(t *testing.T) {
	nums := []int{5}
	got := searchRange(nums, 5)
	if got[0] == -1 || got[1] == -1 {
		t.Fatalf("got: %v, want: %v", got, []int{0, 0})
	}
}
