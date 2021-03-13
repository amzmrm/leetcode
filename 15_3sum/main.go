package leetcode

import "sort"

// 时间复杂度：O(n^2)
// 空间复杂度：O(1)
func ThreeSum(nums []int) [][]int {
	if nums == nil || len(nums) < 3 {
		return nil
	}
	sort.Ints(nums)
	if nums[0] > 0 {
		return nil
	}

	ret := make([][]int, 0)
	sum, left, right := 0, 0, 0
	for i := 0; i < len(nums); i++ {
		for i > 0 && i < len(nums)-1 && nums[i] == nums[i-1] {
			i++
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

func threeSum2(nums []int, target int) [][]int {
	return threeSumTarget(nums, target, false)
}

func threeSumTarget(nums []int, target int, returnIndex bool) [][]int {
	sort.Ints(nums)
	var res [][]int
	for i := 0; i < len(nums); i++ {
		tuples := twoSumTarget(nums, i+1, target-nums[i], returnIndex)
		for _, tuple := range tuples {
			if returnIndex {
				tuple = append(tuple, i)
				res = append(res, tuple)
			} else {
				tuple = append(tuple, nums[i])
				res = append(res, tuple)
			}
		}
		for i < len(nums)-1 && nums[i] == nums[i+1] {
			i++
		}
	}
	return res
}

func twoSumTarget(nums []int, start int, target int, returnIndex bool) [][]int {
	var res [][]int
	sort.Ints(nums)
	left, right, sum := start, len(nums)-1, 0
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
			if returnIndex {
				res = append(res, []int{left, right})
			} else {
				res = append(res, []int{nums[left], nums[right]})
			}
			for left < right && nums[left] == leftElem {
				left++
			}
		}
	}
	return res
}
