package leetcode

func findClosestElements(arr []int, k int, x int) []int {
	left, right, mid := 0, len(arr)-k, 0
	for left < right {
		mid = (left + right) / 2
		if x-arr[mid] > arr[mid+k]-x { // 右边界与 x 的差值的绝对值较小，左边界收缩
			left = mid + 1
		} else {
			// 左边界与 x 的差值的绝对值较小，右边界收缩
			// 或者，左、右边界与 x 的差的绝对值相等，右边界向左收缩
			right = mid
		}
	}
	return arr[left : left+k]
}
