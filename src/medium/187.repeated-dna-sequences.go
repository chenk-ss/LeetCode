package medium

/*
 * @lc app=leetcode id=187 lang=golang
 *
 * [187] Repeated DNA Sequences
 */

// @lc code=start
func findRepeatedDnaSequences(s string) []string {
	res := []string{}
	m := make(map[string]int)
	for i := 0; i < len(s)-9; i++ {
		str := s[i : i+10]
		if m[str] == 1 {
			res = append(res, str)
		}
		m[str]++
	}
	return res
}

// @lc code=end
