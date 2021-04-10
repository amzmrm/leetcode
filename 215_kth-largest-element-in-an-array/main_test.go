package leetcode

import "testing"

func TestFindFindKthLargest(t *testing.T) {
	test := []int{7, 6, 5, 4, 3, 2, 1}
	got := findKthLargest(test, 5)
	t.Log(got)
}
