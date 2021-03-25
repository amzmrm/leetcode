package leetcode

import "math"

func reverse(x int) int {
	var rev int
	for x != 0 {
		pop := x % 10
		x = x / 10
		if rev > math.MaxInt32/10 || (rev == math.MaxInt32 && pop > math.MaxInt32%10) {
			return 0
		} else if rev < math.MinInt32/10 || (rev == math.MaxInt32 && pop < math.MinInt32%10) {
			return 0
		}
		rev = rev*10 + x%10
	}
	return rev
}
