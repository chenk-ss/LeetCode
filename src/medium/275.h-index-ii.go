package medium

/*
 * @lc app=leetcode id=275 lang=golang
 *
 * [275] H-Index II
 */

// @lc code=start
func hIndex275(citations []int) int {
	res := len(citations)
	for _, v := range citations {
		if v < res {
			res--
		}
	}
	return res
}

// @lc code=end
