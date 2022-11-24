package src

/*
 * @lc app=leetcode id=11 lang=golang
 *
 * [11] Container With Most Water
 */

// @lc code=start
func maxArea(height []int) int {
	length := len(height)
	start := 0
	end := length - 1
	maxArea := 0
	for start < end {
		minHeigth := minInt(height[start], height[end])
		newArea := (end - start) * minHeigth
		if newArea > maxArea {
			maxArea = newArea
		}
		if height[start] <= height[end] {
			start++
		} else {
			end--
		}
	}
	return maxArea
}

func minInt(num1 int, num2 int) int {
	if num1 > num2 {
		return num2
	}
	return num1
}

// @lc code=end
