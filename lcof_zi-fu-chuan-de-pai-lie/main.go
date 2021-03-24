package leetcode

func permutation(s string) []string {
	var res []string
	var track string
	used := make([]bool, len(s))

	backtrack(s, &res, &track, &used)

	return res
}

func backtrack(
	s string,
	res *[]string,
	track *string,
	used *[]bool,
) {
	if len(*track) == len(s) {
		str := *track
		*res = append(*res, str)
		return
	}

	set := make(map[uint8]bool)
	for i := 0; i < len(s); i++ {
		if (*used)[i] { // 已经选择过的，跳过
			continue
		}
		if set[s[i]] { // 同一个位置不能选择重复的字符串
			continue
		}
		set[s[i]] = true
		*track = (*track) + string(s[i])
		(*used)[i] = true
		backtrack(s, res, track, used)
		*track = (*track)[:len(*track)-1]
		(*used)[i] = false
	}
}
