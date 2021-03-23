package leetcode

func search(nums []int, target int) int {
	if len(nums) <= 0 {
		return -1
	}

	left, right, mid := 0, len(nums)-1, 0
	for left <= right {
		mid = (right + left) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		}
	}
	return -1
}
