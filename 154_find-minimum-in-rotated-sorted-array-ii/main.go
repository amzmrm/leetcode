package leetcode

func findMin(nums []int) int {
	left, right, mid := 0, len(nums)-1, 0
	for left < right {
		mid = (right + left) / 2
		if nums[mid] == nums[right] {
			right--
		} else if nums[mid] < nums[right] {
			right = mid
		} else if nums[mid] > nums[right] {
			left = mid + 1
		}
	}
	return nums[left]
}
