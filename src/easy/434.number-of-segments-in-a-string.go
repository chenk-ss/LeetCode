package easy

/*
 * @lc app=leetcode id=434 lang=golang
 *
 * [434] Number of Segments in a String
 */

// @lc code=start
func countSegments(s string) int {
	result, flag := 0, true
	for i := range s {
		if s[i] != ' ' && flag {
			result++
		}
		flag = s[i] == ' '
	}
	return result
}

// @lc code=end
