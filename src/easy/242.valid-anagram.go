package easy

/*
 * @lc app=leetcode id=242 lang=golang
 *
 * [242] Valid Anagram
 */

// @lc code=start
func isAnagram(s string, t string) bool {
	list := make([]int, 26)
	for i := range s {
		num := s[i] - 'a'
		list[num]++
	}
	for i := range t {
		num := t[i] - 'a'
		list[num]--
	}
	for i := range list {
		if list[i] != 0 {
			return false
		}
	}
	return true
}

// @lc code=end
