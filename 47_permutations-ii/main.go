package leetcode

import "sort"

func permuteUnique(nums []int) [][]int {
	if len(nums) == 0 {
		return nil
	}

	var res [][]int

	var track []int
	var backtrack func([]int, []int)
	used := make([]bool, len(nums))
	backtrack = func(ns []int, track []int) {
		if len(track) == len(nums) {
			var newTrack []int
			newTrack = append(newTrack, track...)
			res = append(res, newTrack)
			return
		}

		for i := 0; i < len(ns); i++ {
			if used[i] {
				continue
			}

			if i > 0 && ns[i] == ns[i-1] && !used[i-1] {
				continue
			}

			track = append(track, ns[i])
			used[i] = true
			backtrack(ns, track)
			track = track[:len(track)-1]
			used[i] = false
		}
	}

	sort.Ints(nums)
	backtrack(nums, track)

	return res
}
