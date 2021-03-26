package leetcode

func rotate(nums []int, k int) {
	newNums := make([]int, len(nums))
	k = k % len(nums)
	newNums = append(newNums, nums[len(nums)-k:]...)
	newNums = append(newNums, nums[:k]...)
	copy(nums, newNums)
}
