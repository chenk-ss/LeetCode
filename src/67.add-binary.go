package src

/*
 * @lc app=leetcode id=67 lang=golang
 *
 * [67] Add Binary
 */

// @lc code=start
func addBinary(a string, b string) string {
	if len(a) < len(b) {
		a, b = b, a
	}
	indexB := len(b) - 1
	result := make([]byte, len(a))
	var sum, flag byte
	for i := len(a) - 1; i >= 0; i-- {
		sum = a[i] + flag
		if indexB >= 0 {
			sum += b[indexB]
			indexB--
		}
		result[i] = sum%2 + '0'
		flag = sum >> 1 % 2
	}
	if flag != 0 {
		return "1" + string(result)
	}
	return string(result)
}

// @lc code=end
