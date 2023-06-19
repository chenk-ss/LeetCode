package medium

/*
 * @lc app=leetcode id=22 lang=golang
 *
 * [22] Generate Parentheses
 */

// @lc code=start
func generateParenthesis(n int) []string {
	if n == 1 {
		return []string{"()"}
	}
	list := generateParenthesis(n - 1)
	result := []string{}
	m := make(map[string]bool)
	for _, v := range list {
		for i := 0; i <= len(v)/2; i++ {
			val := v[:i] + "()" + v[i:]
			if _, ok := m[val]; !ok {
				result = append(result, val)
				m[val] = true
			}
		}
	}
	return result
}

// @lc code=end
