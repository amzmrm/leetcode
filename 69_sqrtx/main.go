package main

func mySqrt(x int) int {
	left, right, mid, ans, val := 0, x, 0, -1, 0
	for left <= right {
		mid = (right + left) / 2
		val = mid * mid
		if val == x {
			return mid
		} else if val < x {
			ans = mid
			left = mid + 1
		} else if val > x {
			right = mid - 1
		}
	}
	return ans
}
