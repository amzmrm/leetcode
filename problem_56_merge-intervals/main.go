package problem_56_merge_intervals

import "sort"

func Merge(intervals [][]int) [][]int {
	if len(intervals) <= 1 {
		return intervals
	}

	// 先将区间根据起点排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	merged := make([][]int, 0)
	merged = append(merged, intervals[0])
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] > merged[len(merged)-1][1] {
			merged = append(merged, intervals[i])
			continue
		}
		if merged[len(merged)-1][1] < intervals[i][1] {
			merged[len(merged)-1][1] = intervals[i][1]
		}
	}

	return merged
}
