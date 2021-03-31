package leetcode

import "testing"

func TestName(t *testing.T) {
	input := []int{-10, -3, 0, 5, 9}
	got := sortedArrayToBST(input)
	t.Log(got.Val)
}
