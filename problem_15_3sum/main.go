package main

import "sort"

// 时间复杂度：O(n^2)
// 空间复杂度：O(1)
func ThreeSum(nums []int) [][]int {
	ret := make([][]int, 0)

	if nums == nil || len(nums) < 3 {
		return ret
	}

	sort.Ints(nums)

	if nums[0] > 0 {
		return ret
	}

	sum, left, right := 0, 0, 0
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		left, right = i+1, len(nums)-1
		for left < right {
			sum = nums[i] + nums[left] + nums[right]
			if sum == 0 {
				ret = append(ret, []int{nums[i], nums[left], nums[right]})
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				left++
				right--
			} else if sum < 0 {
				left++
			} else {
				right--
			}
		}
	}
	return ret
}
