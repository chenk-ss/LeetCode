package easy

import "math"

/*
 * @lc app=leetcode id=492 lang=golang
 *
 * [492] Construct the Rectangle
 */

// @lc code=start
func constructRectangle(area int) []int {
	sqrt := int(math.Sqrt(float64(area)))
	for area%sqrt != 0 {
		sqrt--
	}
	return []int{area / sqrt, sqrt}
}

// @lc code=end
