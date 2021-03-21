package leetcode

func canJump(nums []int) bool {
	farthest := 0
	for i := 0; i < len(nums); i++ {
		if i <= farthest {
			farthest = max(farthest, i+nums[i])
			if farthest >= len(nums)-1 {
				return true
			}
		}
	}
	return farthest >= len(nums)-1
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
