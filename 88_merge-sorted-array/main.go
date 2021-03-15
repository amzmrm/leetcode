package leetcode

func merge(nums1 []int, m int, nums2 []int, n int) []int {
	i, j := m-1, n-1
	p := m + n - 1
	for i >= 0 && j >= 0 {
		if nums1[i] <= nums2[j] {
			nums1[p] = nums2[j]
			j--
		} else {
			nums1[p] = nums1[i]
			i--
		}
		p--
	}
	for k := j; k >= 0 && p >= 0; k-- {
		nums1[p] = nums2[k]
		p--
	}
	return nums1
}
