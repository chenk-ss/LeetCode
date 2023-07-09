package medium

import "sort"

/*
 * @lc app=leetcode id=452 lang=golang
 *
 * [452] Minimum Number of Arrows to Burst Balloons
 */

// @lc code=start
func findMinArrowShots(points [][]int) int {
	sort.Slice(points, func(i, j int) bool {
		return points[i][1] < points[j][1]
	})
	last := points[0][1]
	count := 1
	for i := 0; i < len(points); i++ {
		if points[i][0] > last {
			count++
			last = points[i][1]
		}
	}
	return count
}

// @lc code=end
