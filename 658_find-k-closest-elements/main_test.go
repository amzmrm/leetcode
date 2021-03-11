package leetcode

import (
	"fmt"
	"testing"
)

func TestName(t *testing.T) {
	nums := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	got := findClosestElements(nums, 2, 6)
	_ = got
	fmt.Println(got)
}
