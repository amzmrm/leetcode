package leetcode

func findAnagrams(s string, p string) []int {
	need, window := make(map[uint8]int, len(p)), make(map[uint8]int, len(s))
	for i := 0; i < len(p); i++ {
		need[p[i]]++
	}

	var res []int
	var left, right, valid int
	var c, d uint8
	for right < len(s) {
		c = s[right]
		right++
		if need[c] > 0 {
			window[c]++
			if window[c] == need[c] {
				valid++
			}
		}

		for right-left >= len(p) {
			if valid == len(need) {
				res = append(res, left)
			}
			d = s[left]
			left++
			if need[d] > 0 {
				if window[d] == need[d] {
					valid--
				}
				window[d]--
			}
		}
	}
	return res
}
