package medium

/*
 * @lc app=leetcode id=397 lang=golang
 *
 * [397] Integer Replacement
 */

// @lc code=start
func integerReplacement(n int) int {
	if n == 1 {
		return 0
	}
	if n == 0 || n == 2 {
		return 1
	}
	if n&1 == 0 {
		return integerReplacement(n/2) + 1
	}
	return 1 + integerReplacementFunc(integerReplacement(n-1), integerReplacement(n+1))
}

func integerReplacementFunc(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// @lc code=end
