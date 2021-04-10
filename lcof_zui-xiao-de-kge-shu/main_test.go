package leetcode

import "testing"

func TestName(t *testing.T) {
	test := []int{0, 1, 2, 3, 4, 5}
	got := getLeastNumbers(test, 4)
	t.Log(got)
}
