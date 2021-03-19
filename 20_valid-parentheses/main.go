package leetcode

func isValid(s string) bool {
	if len(s)%2 != 0 {
		return false
	}

	pairs := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}
	var stack []byte
	for i := 0; i < len(s); i++ {
		if pairs[s[i]] > 0 { // 碰到右括号
			if len(stack) == 0 || stack[len(stack)-1] != pairs[s[i]] {
				return false
			}
			stack = stack[:len(stack)-1]
		} else { // 碰到左括号
			stack = append(stack, s[i])
		}
	}
	return len(stack) == 0
}
