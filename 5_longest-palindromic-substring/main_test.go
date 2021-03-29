package leetcode

import "testing"

func TestLongestPalindrome(t *testing.T) {
	input := "babab"
	got := longestPalindrome2(input)
	t.Log(got)
}
