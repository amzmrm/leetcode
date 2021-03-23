package leetcode

func checkInclusion(s1 string, s2 string) bool {
	need, window := make(map[uint8]uint8, len(s1)), make(map[uint8]uint8, len(s2))
	for i := 0; i < len(s1); i++ {
		need[s1[i]]++
	}

	var left, right, valid int
	var c, d uint8
	for right < len(s2) {
		c = s2[right]
		right++
		if need[c] > 0 {
			window[c]++
			if window[c] == need[c] {
				valid++
			}
		}

		for right-left >= len(s1) {
			if valid == len(need) {
				return true
			}
			d = s2[left]
			left++
			if need[d] > 0 {
				if window[d] == need[d] {
					valid--
				}
				window[d]--
			}
		}
	}
	return false
}
