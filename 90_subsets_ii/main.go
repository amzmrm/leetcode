package leetcode

import "sort"

func subsetsWithDup(nums []int) [][]int {
	var res [][]int
	var track []int

	sort.Ints(nums)
	backtrack(nums, &res, &track, 0)

	return res
}

func backtrack(nums []int, res *[][]int, track *[]int, start int) {
	newTrack := make([]int, len(*track))
	copy(newTrack, *track)
	*res = append(*res, newTrack)

	for i := start; i < len(nums); i++ {
		if i > start && nums[i] == nums[i-1] {
			continue
		}
		*track = append(*track, nums[i])
		backtrack(nums, res, track, i+1)
		*track = (*track)[:len(*track)-1]
	}
}
