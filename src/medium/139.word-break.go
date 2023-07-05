package medium

/*
 * @lc app=leetcode id=139 lang=golang
 *
 * [139] Word Break
 */

// @lc code=start
func wordBreak(s string, wordDict []string) bool {
	list := make([]bool, len(s)+1)
	list[len(s)] = true
	for i := len(s) - 1; i >= 0; i-- {
		for _, str := range wordDict {
			if i+len(str) <= len(s) && s[i:i+len(str)] == str {
				list[i] = list[i+len(str)]
				if list[i] {
					break
				}
			}
		}
	}
	return list[0]
}

// @lc code=end
