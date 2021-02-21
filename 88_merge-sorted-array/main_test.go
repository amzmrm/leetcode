package leetcode

import "testing"

func TestMerge0(t *testing.T) {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	nums2 := []int{2, 5, 6}
	nums1 = Merge0(nums1, 3, nums2, len(nums2))
	t.Logf("%#v", nums1)
}

func TestMerge1(t *testing.T) {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	nums2 := []int{2, 5, 6}
	nums1 = Merge1(nums1, 3, nums2, len(nums2))
	t.Logf("%#v", nums1)
}

func TestMerge2(t *testing.T) {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	nums2 := []int{2, 5, 6}
	nums1 = Merge2(nums1, 3, nums2, len(nums2))
	t.Logf("%#v", nums1)
}
