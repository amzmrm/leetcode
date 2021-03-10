package leetcode

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	if len(nums) == 1 {
		if nums[0] == target {
			return 0
		}
		return -1
	}

	left, right, mid := 0, len(nums)-1, 0
	for left <= right {
		mid = (right + left) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] >= nums[left] { // [0, mid]这一段是有序的
			if nums[left] <= target && target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else if nums[mid] <= nums[right] { // [mid, last]这一段是有序的
			if nums[mid] < target && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}

	return -1
}
