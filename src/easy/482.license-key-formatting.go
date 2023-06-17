package easy

import "strings"

/*
 * @lc app=leetcode id=482 lang=golang
 *
 * [482] License Key Formatting
 */

// @lc code=start
func licenseKeyFormatting(s string, k int) string {
	s = strings.Replace(s, "-", "", -1)
	result := ""
	count := 0
	for i := len(s) - 1; i >= 0; i-- {
		if count == k {
			result = "-" + result
			count = 0
		}
		count++
		if s[i]-'a' >= 0 && s[i]-'a' <= 25 {
			result = string(s[i]-'a'+'A') + result
		} else {
			result = string(s[i]) + result
		}
	}
	return result
}

// @lc code=end
