package leetcode

import "sort"

func twoSum(nums []int, target int) []int {
	m := make(map[int]int, 0)
	for i := 0; i < len(nums); i++ {
		complement := target - nums[i]
		if cpm, ok := m[complement]; ok {
			return []int{cpm, i}
		}
		m[nums[i]] = i
	}
	return []int{}
}

func twoSum2(nums []int, target int) []int {
	res := twoSumTarget(nums, target)
	if len(res) != 0 {
		return res[0]
	}
	return nil
}

func twoSumTarget(nums []int, target int) [][]int {
	var res [][]int
	sort.Ints(nums)
	left, right, sum := 0, len(nums)-1, 0
	leftElem, rightElem := 0, 0
	for left < right {
		leftElem, rightElem = nums[left], nums[right]
		sum = leftElem + rightElem
		if sum < target {
			for left < right && nums[left] == leftElem {
				left++
			}
		} else if sum > target {
			for left < right && nums[right] == rightElem {
				right--
			}
		} else {
			res = append(res, []int{left, right})
			for left < right && nums[left] == leftElem {
				left++
			}
			for left < right && nums[right] == rightElem {
				right--
			}
		}
	}
	return res
}
