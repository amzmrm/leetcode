package leetcode

func jump(nums []int) int {
	end, maxPos, steps := 0, 0, 0
	for i := 0; i < len(nums)-1; i++ {
		maxPos = max(nums[i]+i, maxPos)
		if end == i {
			steps++
			end = maxPos
		}
	}
	return steps
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
