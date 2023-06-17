package medium

/*
 * @lc app=leetcode id=5 lang=golang
 *
 * [5] Longest Palindromic Substring
 */

// @lc code=start
func stringLongestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}
	start := 0
	end := 0
	i := 0
	for i < len(s) {
		L1 := check(s, i, i)
		L2 := check(s, i, i+1)
		maxLen := max(L1, L2)
		if maxLen > end-start {
			start = i - (maxLen-1)/2
			end = i + maxLen/2
		}
		i++
	}
	return s[start : end+1]
}

func max(num1 int, num2 int) int {
	if num1 > num2 {
		return num1
	}
	return num2
}

func check(s string, left int, right int) int {
	L := left
	R := right
	for L >= 0 && R < len(s) && s[L:L+1] == s[R:R+1] {
		L--
		R++
	}
	return R - L - 1
}

// @lc code=end
