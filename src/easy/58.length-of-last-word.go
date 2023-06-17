package easy

/*
 * @lc app=leetcode id=58 lang=golang
 *
 * [58] Length of Last Word
 */

// @lc code=start
func lengthOfLastWord(s string) int {
	flag := true
	length := 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i:i+1] == " " {
			if !flag {
				return length
			}
		} else {
			length++
			flag = false
		}
	}
	return length
}

// @lc code=end
