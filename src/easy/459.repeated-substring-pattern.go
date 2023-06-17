package easy

import "strings"

/*
 * @lc app=leetcode id=459 lang=golang
 *
 * [459] Repeated Substring Pattern
 */

// @lc code=start
func repeatedSubstringPattern(s string) bool {
	for i := 0; i < len(s)/2; i++ {
		str := s[:i+1]
		s2 := strings.Replace(s, str, "", -1)
		if s2 == "" {
			return true
		}
	}
	return false
}

func repeatedSubstringPattern2(s string) bool {
	s1 := s[1:] + s[0:len(s)-1]
	return strings.Contains(s1, s)
}

// @lc code=end
