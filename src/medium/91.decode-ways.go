package medium

/*
 * @lc app=leetcode id=91 lang=golang
 *
 * [91] Decode Ways
 */

// @lc code=start
func numDecodings(s string) int {
	if s[0] == '0' {
		return 0
	}
	list := make([]int, len(s)+1)
	list[0], list[1] = 1, 1
	for i := 2; i <= len(s); i++ {
		str := s[i-2 : i]
		if str >= "10" && str <= "26" {
			list[i] = list[i-2]
		}
		if s[i-1] != '0' {
			list[i] += list[i-1]
		}
	}
	return list[len(s)]
}

// @lc code=end
