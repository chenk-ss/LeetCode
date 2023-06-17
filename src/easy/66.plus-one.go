package easy

/*
 * @lc app=leetcode id=66 lang=golang
 *
 * [66] Plus One
 */

// @lc code=start
func plusOne(digits []int) []int {
	length := len(digits)
	for i := length - 1; i >= 0; i-- {
		if digits[i] != 9 {
			digits[i]++
			return digits
		}
		digits[i] = 0
	}
	digits = make([]int, length+1)
	digits[0] = 1
	return digits
}

// @lc code=end
