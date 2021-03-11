package leetcode

func findPeakElement(nums []int) int {
	l, r, mid := 0, len(nums)-1, 0
	for l < r {
		mid = (l + r) / 2
		if nums[mid] > nums[mid+1] { // 该元素恰好位于降序序列或者一个局部下降坡度中，峰值会在本元素的左边。
			r = mid
		} else { // 该元素恰好位于升序序列或者一个局部上升坡度中，峰值会在本元素的右边。
			l = mid + 1
		}
	}
	return l
}
