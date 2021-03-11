package leetcode

/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

func firstBadVersion(n int) int {
	left, right, mid := 1, n, 0

	var ok bool
	for left < right {
		mid = (left + right) / 2
		ok = isBadVersion(mid)
		if ok {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return left
}
