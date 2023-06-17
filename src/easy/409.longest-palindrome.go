package easy

/*
 * @lc app=leetcode id=409 lang=golang
 *
 * [409] Longest Palindrome
 */

// @lc code=start
func longestPalindrome(s string) int {
	list := make([]int, 52)
	for i := range s {
		if s[i]-'A' > 25 {
			list[s[i]-'a'+26]++
		} else {
			list[s[i]-'A']++
		}
	}
	sum := 0
	flag := false
	for i := range list {
		if !flag && list[i]%2 != 0 {
			flag = true
		}
		if list[i] > 1 {
			sum += list[i] / 2 * 2
		}
	}
	if flag {
		sum++
	}
	return sum
}

// @lc code=end
