package leetcode

import "testing"

func TestMerge(t *testing.T) {
	nums1 := []int{1, 2, 0, 0, 0, 0, 0, 0}
	nums2 := []int{3, 4, 6, 7, 8, 9}
	nums1 = merge(nums1, 2, nums2, len(nums2))
	t.Logf("%#v", nums1)
}
