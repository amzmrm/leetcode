package main

func generateParenthesis(n int) []string {
	if n == 0 {
		return nil
	}

	// 记录所有合法的括号组合
	var res []string
	// 回溯过程中的路径
	var track string
	backtrack(n, n, &res, &track)

	return res
}

// left, right是指可用的左括号和右括号数量
func backtrack(left int, right int, res *[]string, track *string) {
	// 左括号应该消耗得比右括号快，否则就是不合法了
	if left > right {
		return
	}
	// 数量小于 0 肯定是不合法的
	if left < 0 || right < 0 {
		return
	}
	// 当所有括号都恰好用完时，得到一个合法的括号组合
	if left == 0 && right == 0 {
		s := []rune(*track)
		*res = append(*res, string(s))
		return
	}

	// 尝试放一个左括号
	*track += "("
	backtrack(left-1, right, res, track)
	*track = (*track)[:len(*track)-1]

	// 尝试放一个右括号
	*track += ")"
	backtrack(left, right-1, res, track)
	*track = (*track)[:len(*track)-1]
}
