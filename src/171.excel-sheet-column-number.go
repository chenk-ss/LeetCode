package src

/*
 * @lc app=leetcode id=171 lang=golang
 *
 * [171] Excel Sheet Column Number
 */

// @lc code=start
func titleToNumber(columnTitle string) int {
	result := 0
	for i := 0; i < len(columnTitle); i++ {
		a := int(columnTitle[i]-'A') + 1
		result = result*26 + a
	}
	return result
}

// @lc code=end
