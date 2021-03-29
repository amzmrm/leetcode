package leetcode

func longestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}
	var start, end int
	for i := 0; i < len(s); i++ {
		oddLeft, oddRight := expandAroundCenter(s, i, i)
		evenLeft, evenRight := expandAroundCenter(s, i, i+1)
		if oddRight-oddLeft > end-start {
			start, end = oddLeft, oddRight
		}
		if evenRight-evenLeft > end-start {
			start, end = evenLeft, evenRight
		}
	}
	return s[start : end+1]
}
func expandAroundCenter(s string, left int, right int) (int, int) {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left, right = left-1, right+1
	}
	return left + 1, right - 1
}

// 当子串length>2时，状态转移方程：P(i,j)=P(i+1,j-1) && Si == Sj
// 当子串length为1或者2时，状态转移方程：P(i,i) = true或者P(i,i+1)=(Si==Si+1)
func longestPalindrome2(s string) string {
	n := len(s)
	res := ""
	dp := make([][]bool, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]bool, n)
	}
	var end int
	for j := 0; j < n; j++ { // 由方程得知，dp[i][j]的值依赖左下角的值（即dp[i+1][j-1]），所以按照：1)先升序填列、2)先升序填行
		for i := 0; i+j < n; i++ { // 由于i,j的关系是i<=j，因此只需填矩阵的右上部分即可，所以i+j<n
			end = i + j
			if j == 0 {
				dp[i][i] = true // 填对角线处都为true（P(i,i) = true）
			} else if j == 1 {
				if s[i] == s[end] { // 字串length==2，两个字符相等的话，就是回文（(i,i+1)=(Si==Si+1)）
					dp[i][end] = true
				}
			} else {
				if s[i] == s[end] { // P(i,j)=P(i+1,j-1) && Si == Sj
					dp[i][end] = dp[i+1][end-1]
				}
			}
			if dp[i][end] && j+1 > len(res) { // j是包含的，所以j+1
				res = s[i : end+1]
			}
		}
	}
	return res
}
