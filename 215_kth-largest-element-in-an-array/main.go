package leetcode

func findKthLargest(nums []int, k int) int {
	partitionArr(nums, 0, len(nums)-1, len(nums)-k)
	return nums[len(nums)-k]
}

func partitionArr(nums []int, left int, right int, index int) {
	if left >= right {
		return
	}
	pos := partition(nums, left, right)
	if pos == index {
		return
	} else if pos < index {
		partitionArr(nums, pos+1, right, index)
	} else {
		partitionArr(nums, left, pos-1, index)
	}
}

func partition(nums []int, left, right int) int {
	i, j, p := left, right, right
	for i != j {
		for nums[i] < nums[p] {
			i++
		}
		for i < j && nums[j] >= nums[p] {
			j--
		}
		if i != j {
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	nums[i], nums[p] = nums[p], nums[i]
	return i
}
