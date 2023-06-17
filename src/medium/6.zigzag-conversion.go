package medium

/*
 * @lc app=leetcode id=6 lang=golang
 *
 * [6] Zigzag Conversion
 */

// @lc code=start
func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	len := len(s)
	data := make([]byte, len)

	interval := numRows*2 - 2
	index := 0
	for i := 0; i < numRows; i++ {
		for j := 0; j+i < len; j += interval {
			data[index] = s[j+i]
			index++
			if i != 0 && i != numRows-1 && j+interval-i < len {
				data[index] = s[j+interval-i]
				index++
			}
		}
	}
	return string(data)
}

// @lc code=end
