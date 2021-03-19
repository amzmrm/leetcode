package main

import "sort"

// 1. 首先为所有区间按照end从小到大排序
// 2. 排序后的区间列表里，每次都选择end最小的
func eraseOverlapIntervals(intervals [][]int) int {
	if len(intervals) <= 1 {
		return 0
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})

	var eraseCount int
	var end int
	for i := range intervals {
		if i == 0 {
			end = intervals[i][1]
			continue
		}
		if intervals[i][0] < end {
			eraseCount++
			continue
		}
		end = intervals[i][1]
	}

	return eraseCount
}
