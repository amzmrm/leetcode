package leetcode

func sortArray(nums []int) []int {
	qsort(nums, 0, len(nums)-1)
	return nums
}

func qsort(nums []int, left, right int) {
	if left >= right {
		// 当数组只有一个元素的时候退出递归
		// 只有一个元素qsort(nums, left, pIdx-1)
		// 和qsort(nums, pIdx+1, right)会导致数组越界
		return
	}
	pIdx := partition(nums, left, right)
	qsort(nums, left, pIdx-1)  // 递归处理停止位置左边的数列
	qsort(nums, pIdx+1, right) // 递归处理停止位置右边的数列
}

func partition(nums []int, left, right int) int {
	i, j, pivot := left, right, nums[right] // 以最右边的数字为基准
	for i != j {
		for nums[i] < pivot { // 从左边开始，直到找到大于等于基准的数的位置
			i++
		}
		for nums[j] >= pivot && j > i { // 从右边开始，直到找到小于基准的数的位置
			j--
		}
		// 左右标记都停下来了，如果不是在同一个位置停下来，就交换左右位置的数字，
		// 从而达到小于基准的数放在左边，大于或等于基准的数放在右边
		if i != j {
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	// 左右标记相遇，本轮对比结束，交换结束位置和基准的数，从而达到基准归位的目的
	// 从下面这行代码就可以看出快排的不稳定：
	// 当i=3是，nums[i]==8==pivot，但是pivot和arr[i]却需要对换
	nums[i], nums[right] = pivot, nums[i]
	return i
}
