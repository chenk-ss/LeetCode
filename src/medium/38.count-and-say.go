package medium

import "strconv"

/*
 * @lc app=leetcode id=38 lang=golang
 *
 * [38] Count and Say
 */

// @lc code=start
func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}
	str := countAndSay(n - 1)
	count := 0
	s := byte('0')
	result := ""
	for i := range str {
		if str[i] == s {
			count++
		} else {
			if count != 0 {
				result += strconv.Itoa(count) + string(s)
			}
			count = 1
			s = str[i]
		}
	}
	result += strconv.Itoa(count) + string(s)
	return result
}

// @lc code=end
