package medium

import "sort"

/*
 * @lc app=leetcode id=57 lang=golang
 *
 * [57] Insert Interval
 */

// @lc code=start
func insert(intervals [][]int, newInterval []int) [][]int {
	intervals = append(intervals, newInterval)
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
