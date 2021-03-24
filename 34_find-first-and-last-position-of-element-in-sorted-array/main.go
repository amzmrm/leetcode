package leetcode

func searchRange(nums []int, target int) []int {
	notFound := []int{-1, -1}
	if len(nums) == 0 {
		return notFound
	}

	l := binarySearch(nums, target, true)
	r := binarySearch(nums, target, false) - 1
	if l <= r && r < len(nums) && nums[l] == target && nums[r] == target {
		return []int{l, r}
	}

	return notFound
}

func binarySearch(nums []int, target int, first bool) int {
	l, r, mid, ans := 0, len(nums)-1, 0, len(nums)
	if first {
		for l <= r { // 找第一个大于或等于target的数的下标
			mid = (l + r) / 2
			if nums[mid] >= target {
				r = mid - 1
				ans = mid
			} else {
				l = mid + 1
			}
		}
	} else {
		for l <= r { // 找第一个大于target的数的下标
			mid = (l + r) / 2
			if nums[mid] > target {
				r = mid - 1
				ans = mid
			} else {
				l = mid + 1
			}
		}
	}
	return ans
}
