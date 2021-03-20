package leetcode

import "sort"

func findMinArrowShots(points [][]int) int {
	if len(points) < 1 {
		return 0
	}

	sort.Slice(points, func(i, j int) bool {
		return points[i][1] < points[j][1]
	})

	var arrows int
	var end int
	for i := range points {
		if i == 0 {
			arrows++
			end = points[i][1]
			continue
		}
		if points[i][0] <= end {
			continue
		}
		end = points[i][1]
		arrows++
	}

	return arrows
}
