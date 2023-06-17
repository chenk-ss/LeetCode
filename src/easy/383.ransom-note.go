package easy

/*
 * @lc app=leetcode id=383 lang=golang
 *
 * [383] Ransom Note
 */

// @lc code=start
func canConstruct(ransomNote string, magazine string) bool {
	list := make([]int, 26)
	for i := range ransomNote {
		list[ransomNote[i]-'a']++
	}
	for i := range magazine {
		list[magazine[i]-'a']--
	}
	for i := range list {
		if list[i] > 0 {
			return false
		}
	}
	return true
}

// @lc code=end
