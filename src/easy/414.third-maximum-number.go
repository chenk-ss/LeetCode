package easy

import "math"

/*
 * @lc app=leetcode id=414 lang=golang
 *
 * [414] Third Maximum Number
 */

// @lc code=start
func thirdMax(nums []int) int {
	first, second, third := math.MinInt64, math.MinInt64, math.MinInt64
	for _, val := range nums {
		if val == first || val == second || val == third {
			continue
		}
		if val > first {
			first, second, third = val, first, second
		} else if val > second {
			second, third = val, second
		} else if val > third {
			third = val
		}
	}
	if third == math.MinInt64 {
		return first
	}
	return third
}

// @lc code=end
