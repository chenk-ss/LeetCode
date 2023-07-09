package medium

import "sort"

/*
 * @lc app=leetcode id=274 lang=golang
 *
 * [274] H-Index
 */

// @lc code=start
func hIndex274(citations []int) int {
	res := len(citations)
	sort.Ints(citations)
	for _, v := range citations {
		if v < res {
			res--
		}
	}
	return res
}

// @lc code=end
