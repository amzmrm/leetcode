package leetcode

import "sort"

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	var res [][]int
	for i := 0; i < len(nums); i++ {
		triples := threeSumTarget(nums, i+1, target-nums[i], false)
		for _, triple := range triples {
			triple = append(triple, nums[i])
			res = append(res, triple)
		}
		for i < len(nums)-1 && nums[i] == nums[i+1] {
			i++
		}
	}
	return res
}

func threeSumTarget(nums []int, start int, target int, returnIndex bool) [][]int {
	sort.Ints(nums)
	var res [][]int
	for i := start; i < len(nums); i++ {
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

func nSum(nums []int, n int, start int, target int) [][]int {
	var res [][]int
	if n < 2 || len(nums) < n { // 至少是 2Sum，且数组大小不应该小于 n
		return res
	}
	if n == 2 { // 2Sum 是 base case
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
				res = append(res, []int{nums[left], nums[right]})
				for left < right && nums[left] == leftElem {
					left++
				}
				for left < right && nums[right] == rightElem {
					right--
				}
			}
		}
	} else {
		for i := start; i < len(nums); i++ {
			sub := nSum(nums, n-1, i+1, target-nums[i])
			for _, ints := range sub {
				ints = append(ints, nums[i])
				res = append(res, ints)
			}
			for i < len(nums)-1 && nums[i] == nums[i+1] {
				i++
			}
		}
	}
	return res
}
