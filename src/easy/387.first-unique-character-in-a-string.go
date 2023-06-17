package easy

/*
 * @lc app=leetcode id=387 lang=golang
 *
 * [387] First Unique Character in a String
 */

// @lc code=start
func firstUniqChar(s string) int {
	list := make([]int, 26)
	for i := range s {
		list[s[i]-'a']++
	}
	for i := range s {
		if list[s[i]-'a'] == 1 {
			return i
		}
	}
	return -1
}

// @lc code=end
