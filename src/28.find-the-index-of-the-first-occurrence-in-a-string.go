package src

/*
 * @lc app=leetcode id=28 lang=golang
 *
 * [28] Find the Index of the First Occurrence in a String
 */

// @lc code=start
func strStr(haystack string, needle string) int {
	nLen := len(needle)
	for i := 0; i <= len(haystack)-nLen; i++ {
		if haystack[i:i+nLen] == needle {
			return i
		}
	}
	return -1
}

// @lc code=end
