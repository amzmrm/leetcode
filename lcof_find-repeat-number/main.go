package leetcode

func findRepeatNumber(nums []int) int {
	vals := make(map[int]struct{})
	for i := 0; i < len(nums); i++ {
		if _, ok := vals[nums[i]]; ok {
			return nums[i]
		} else {
			vals[nums[i]] = struct{}{}
		}
	}
	return -1
}
