package easy

/*
 * @lc app=leetcode id=278 lang=golang
 *
 * [278] First Bad Version
 */

// @lc code=start
/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

func firstBadVersion(n int) int {
	if n == 1 {
		return 1
	}
	left, right, mid := 1, n, 0
	for left <= right {
		mid = (left + right) / 2
		if isBadVersion(mid) {
			right = mid - 1
		} else if isBadVersion(mid + 1) {
			return mid + 1
		} else {
			left = mid + 1
		}
	}
	return left
}

func isBadVersion(version int) bool {
	return true
}

func firstBadVersion2(n int) int {
	if n == 1 {
		return 1
	}
	for n > 0 && isBadVersion(n) {
		n--
	}
	return n + 1
}

// @lc code=end
