package easy

/*
 * @lc app=leetcode id=389 lang=golang
 *
 * [389] Find the Difference
 */

// @lc code=start
func findTheDifference(s string, t string) byte {
	list := make([]int, 26)
	for i := range s {
		list[s[i]-'a']++
	}
	for i := range t {
		list[t[i]-'a']--
	}
	for i := range list {
		if list[i] < 0 {
			return byte('a' + i)
		}
	}
	return 'a'
}

// @lc code=end
