package medium

/*
 * @lc app=leetcode id=131 lang=golang
 *
 * [131] Palindrome Partitioning
 */

// @lc code=start
func partition(s string) [][]string {
	res := [][]string{}
	partition131Func(s, &res, []string{})
	return res
}

func partition131Func(s string, res *[][]string, list []string) {
	if len(s) == 0 {
		if len(list) != 0 {
			*res = append(*res, append([]string{}, list...))
		}
		return
	}
	for i := 1; i <= len(s); i++ {
		if partition131CHeckPalindorm(s[:i]) {
			partition131Func(s[i:], res, append(list, s[:i]))
		}
	}
}

func partition131CHeckPalindorm(s string) bool {
	if len(s) < 2 {
		return true
	}
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}

// @lc code=end
