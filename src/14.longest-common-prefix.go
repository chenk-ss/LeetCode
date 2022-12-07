package src

/*
 * @lc app=leetcode id=14 lang=golang
 *
 * [14] Longest Common Prefix
 */

// @lc code=start
func longestCommonPrefix(strs []string) string {
	data0 := strs[0]
	index := 0
	for index < len(data0) {
		str0 := data0[index]
		for i := 1; i < len(strs); i++ {
			if index >= len(strs[i]) || str0 != strs[i][index] {
				return data0[0:index]
			}
		}
		index++
	}
	return data0[0:index]
}

// @lc code=end
