package medium

import "math"

/*
 * @lc app=leetcode id=29 lang=golang
 *
 * [29] Divide Two Integers
 */

// @lc code=start
func divide(dividend int, divisor int) int {
	if dividend == math.MinInt32 && divisor == -1 {
		return math.MaxInt32
	}
	divd, divr := dividend, divisor
	if dividend < 0 {
		divd = -dividend
	}
	if divisor < 0 {
		divr = -divisor
	}
	count := 0
	for divd >= divr {
		sub := divr
		add := 1
		for divd >= sub<<1 {
			sub <<= 1
			add <<= 1
		}
		divd -= sub
		count += add
	}
	flag := (dividend < 0) != (divisor < 0)
	if flag {
		return -count
	}
	return count
}

// @lc code=end
