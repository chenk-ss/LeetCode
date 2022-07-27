package src

import (
	"strings"
)

/*
 * @lc app=leetcode id=3 lang=golang
 *
 * [3] Longest Substring Without Repeating Characters
 */

// @lc code=start
func lengthOfLongestSubstring(s string) int {
	if len(s) <= 1 {
		return len(s)
	}
	i := 0
	j := 1
	max := 0
	for j < len(s) {
		str := s[i:j]
		for strings.Contains(str, s[j:j+1]) && i < j {
			i++
			str = s[i:j]
		}
		if len(str)+1 > max {
			max = len(str) + 1
		}
		j++
	}
	return max
}

// @lc code=end
