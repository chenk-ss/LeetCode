package easy

/*
 * @lc app=leetcode id=258 lang=golang
 *
 * [258] Add Digits
 */

// @lc code=start
func addDigits(num int) int {
	if num == 0 {
		return 0
	}
	if num%9 == 0 {
		return 9
	}
	return num % 9
}

func addDigits2(num int) int {
	sum := 0
	for num > 0 {
		sum += num % 10
		num /= 10
	}
	for sum >= 10 {
		sum = addDigits(sum)
	}
	return sum
}

// @lc code=end
