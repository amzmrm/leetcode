package leetcode

import "testing"

func TestName(t *testing.T) {
	twoSum := Constructor()
	twoSum.Add(0)
	twoSum.Add(-1)
	twoSum.Add(1)
	ok := twoSum.Find(0)
	t.Log(ok)
}
