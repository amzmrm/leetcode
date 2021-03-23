package leetcode

func lengthOfLongestSubstring(s string) int {
	window := make(map[uint8]int, len(s))
	left, right, res := 0, 0, 0
	var c, d uint8
	for right < len(s) {
		c = s[right]
		right++
		window[c]++
		for window[c] > 1 {
			d = s[left]
			left++
			window[d]--
		}
		res = max(res, right-left)
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
