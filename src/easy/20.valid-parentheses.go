package easy

/*
 * @lc app=leetcode id=20 lang=golang
 *
 * [20] Valid Parentheses
 */

// @lc code=start
func isValid(s string) bool {
	stack := []rune{}

	for _, char := range s {
		switch char {
		case '(':
			stack = append(stack, ')')
		case '[':
			stack = append(stack, ']')
		case '{':
			stack = append(stack, '}')
		default:
			length := len(stack)
			if length == 0 {
				return false
			} else {
				pop := stack[length-1]
				stack = stack[:length-1]
				if pop != char {
					return false
				}
			}
		}
	}
	return len(stack) == 0
}

// @lc code=end
