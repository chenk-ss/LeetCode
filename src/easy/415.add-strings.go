package easy

/*
 * @lc app=leetcode id=415 lang=golang
 *
 * [415] Add Strings
 */

// @lc code=start
func addStrings(num1 string, num2 string) string {
	sum, flag := 0, 0
	list := []byte{}
	for i, j := len(num1)-1, len(num2)-1; i >= 0 || j >= 0; i, j = i-1, j-1 {
		sum = flag
		if i >= 0 {
			sum += int(num1[i] - '0')
		}
		if j >= 0 {
			sum += int(num2[j] - '0')
		}
		flag = sum / 10
		list = append(list, byte(sum%10)+'0')
	}
	if flag == 1 {
		list = append(list, '1')
	}
	for i, j := 0, len(list)-1; i < j; i, j = i+1, j-1 {
		list[i], list[j] = list[j], list[i]
	}

	return string(list)
}

// @lc code=end
