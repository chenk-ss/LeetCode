package easy

/*
 * @lc app=leetcode id=405 lang=golang
 *
 * [405] Convert a Number to Hexadecimal
 */

// @lc code=start
func toHex(num int) string {
	if num == 0 {
		return "0"
	}
	index, sb, str := 7, make([]byte, 8), "0123456789abcdef"
	for i := 7; i >= 0; i-- {
		b := num & 15
		if b > 0 {
			index = i
		}
		sb[i] = str[b]
		num >>= 4
	}
	return string(sb[index:])
}

// @lc code=end
