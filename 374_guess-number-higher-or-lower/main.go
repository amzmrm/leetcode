package main

/**
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is lower than the guess number
 *			      1 if num is higher than the guess number
 *               otherwise return 0
 * func guess(num int) int;
 */

func guessNumber(n int) int {
	low, high := 0, n
	for low <= high {
		mid := low + (high-low)/2
		res := guess(mid)
		if res == 0 {
			return mid
		} else if res == 1 { // the number is higher than mid
			low = mid + 1
		} else if res == -1 { // the number is lower than mid
			high = mid - 1
		}
	}
	return -1
}
