package leetcode

import "math"

// 时间复杂度：O(n)
// 空间复杂度：O(1)
func firstMissingPositive(nums []int) int {
	length := len(nums)

	// 空切片，返回1
	if length == 0 {
		return 1
	}

	count := 0
	// 检查 1 是否存在于数组中
	// O(n)
	for _, num := range nums {
		if num == 1 {
			count++
		}
	}

	// 如果nums没有1，那么缺失的第一个正整数就是1
	if count == 0 {
		return 1
	}

	// nums包含有1，但是length为1，那么缺失的第一个正整数就是2
	if length == 1 {
		return 2
	}

	// 将负数和大于length的数替换为1
	// O(n)
	for i, num := range nums {
		if num <= 0 || num > length {
			nums[i] = 1
		}
	}

	// 遍历数组。当读到数字a时，替换第a个元素的符号
	// O(n)
	a := 0
	for i := 0; i < length; i++ {
		a = int(math.Abs(float64(nums[i])))
		if a == length {
			nums[0] = -int(math.Abs(float64(nums[0])))
		} else {
			nums[a] = -int(math.Abs(float64(nums[a])))
		}
	}

	// 再次遍历数组。返回第一个正数元素的下标
	// O(n)
	for i := 1; i < length; i++ {
		if nums[i] > 0 {
			return i
		}
	}

	// 如果nums[0] > 0，则返回length。这种情况是缺失的第一个正整数就是数组的长度
	if nums[0] > 0 {
		return length
	}

	// 如果之前的步骤中没有发现nums中有正数元素，则返回n + 1
	// 这种情况是数组刚好包含了length个不同的正整数，缺失的第一个正整数就是length+1
	return length + 1
}
