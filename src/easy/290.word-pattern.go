package easy

import "strings"

/*
 * @lc app=leetcode id=290 lang=golang
 *
 * [290] Word Pattern
 */

// @lc code=start
func wordPattern(pattern string, s string) bool {
	list := make([]string, 26)
	sList := strings.Split(s, " ")
	if len(pattern) != len(sList) {
		return false
	}
	m := make(map[string]byte)
	for i := range pattern {
		ch := int(pattern[i] - 'a')
		if list[ch] == "" {
			if _, ok := m[sList[i]]; ok {
				return false
			}
			list[ch] = sList[i]
			m[sList[i]] = byte(pattern[i])
		} else if list[ch] != sList[i] {
			return false
		}
	}
	return true
}

// @lc code=end
