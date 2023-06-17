package easy

/*
 * @lc app=leetcode id=168 lang=golang
 *
 * [168] Excel Sheet Column Title
 */

// @lc code=start
func convertToTitle(columnNumber int) string {
	list := []byte{}
	for columnNumber > 0 {
		columnNumber--
		left := columnNumber % 26
		columnNumber /= 26
		s := byte('A' + left)
		list = append(list, s)
	}
	for i, j := 0, len(list)-1; i < j; i, j = i+1, j-1 {
		list[i], list[j] = list[j], list[i]
	}
	return string(list[:])
}

// @lc code=end
