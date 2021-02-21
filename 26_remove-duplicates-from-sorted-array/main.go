package leetcode

func RemoveDuplicates(nums []int) int {
	originalLen := len(nums)
	if originalLen <= 1 {
		return originalLen
	}

	prev := nums[0]

	// 记录删除的数字个数
	j := 0
	for i := 1; i < len(nums); {
		if i+j >= originalLen {
			break
		}

		if nums[i] == prev {
			nums = append(nums[:i], nums[i+1:]...)
			j++
		} else {
			prev = nums[i]
			i++
		}
	}
	return len(nums)
}

func RemoveDuplicates2(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}

	i := 0
	for j := 1; j < len(nums); j++ {
		if nums[i] != nums[j] {
			i++
			nums[i] = nums[j]
		}
	}
	return i + 1
}
