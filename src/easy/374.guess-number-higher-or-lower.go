package easy

/*
 * @lc app=leetcode id=374 lang=golang
 *
 * [374] Guess Number Higher or Lower
 */

// @lc code=start
/**
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is higher than the picked number
 *			      1 if num is lower than the picked number
 *               otherwise return 0
 * func guess(num int) int;
 */

func guessNumber(n int) int {
	left, mid, right := 1, 1, n
	for left <= right {
		mid = (left + right) / 2
		guessNum := guess(mid)
		if guessNum == 0 {
			return mid
		} else if guessNum == -1 {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return left
}
func guess(num int) int {
	return 0
}

// @lc code=end
