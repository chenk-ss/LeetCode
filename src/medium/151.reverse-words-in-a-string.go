package medium

/*
 * @lc app=leetcode id=151 lang=golang
 *
 * [151] Reverse Words in a String
 */

// @lc code=start
func reverseWords(s string) string {
	res := ""
	left, right := len(s)-1, len(s)-1
	for left >= 0 && right >= 0 {
		if s[right] == ' ' {
			right--
			continue
		}
		left = right
		for left > 0 && s[left-1] != ' ' {
			left--
			continue
		}
		res += " " + s[left:right+1]
		right = left - 1
	}
	return res[1:]
}

// @lc code=end
