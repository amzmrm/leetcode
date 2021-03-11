package leetcode

func peakIndexInMountainArray(arr []int) int {
	l, r, mid := 0, len(arr)-1, 0
	for l < r {
		mid = (l + r) / 2
		if arr[mid] > arr[mid+1] {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return l
}
