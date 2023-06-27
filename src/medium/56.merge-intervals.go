package medium

import "sort"

/*
 * @lc app=leetcode id=56 lang=golang
 *
 * [56] Merge Intervals
 */

// @lc code=start
func merge(intervals [][]int) [][]int {
	sort.SliceStable(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	result := [][]int{}
	prev := intervals[0]
	for _, v := range intervals[1:] {
		if prev[1] < v[0] {
			result = append(result, prev)
			prev = v
		} else if prev[1] < v[1] {
			prev[1] = v[1]
		}
	}
	result = append(result, prev)
	return result
}

// @lc code=end
