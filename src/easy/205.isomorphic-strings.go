package easy

/*
 * @lc app=leetcode id=205 lang=golang
 *
 * [205] Isomorphic Strings
 */

// @lc code=start
func isIsomorphic(s string, t string) bool {
	ss, tt := make(map[byte]int), make(map[byte]int)
	for i := range s {
		if ss[s[i]] != tt[t[i]] {
			return false
		} else {
			ss[s[i]] = i + 1
			tt[t[i]] = i + 1
		}
	}
	return true
}

// @lc code=end
