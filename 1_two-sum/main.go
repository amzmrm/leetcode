package leetcode

import "sort"

func twoSum(nums []int, target int) []int {
	hashTable := map[int]int{}
	for i, num := range nums {
		if idx, ok := hashTable[target-num]; ok {
			return []int{idx, i}
		}
		hashTable[num] = i
	}
	return nil
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
