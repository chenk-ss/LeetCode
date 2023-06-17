package medium

/*
 * @lc app=leetcode id=12 lang=golang
 *
 * [12] Integer to Roman
 */

// @lc code=start
func intToRoman(num int) string {
	result := ""
	if num < 1 && num > 3999 {
		return result
	}
	i := 0
	var numArray = []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	var strArray = []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	for num > 0 {
		for num >= numArray[i] {
			result += strArray[i]
			num -= numArray[i]
		}
		i++
	}
	return result
}

func intToRoman2(num int) string {
	result := ""
	if num < 1 && num > 3999 {
		return result
	}
	if num >= 1000 {
		i := num / 1000
		result = calc(i, result, "M")
		num = num % 1000
	}
	num, result = calcStr(num, result, 100, 500, 1000, "C", "D", "M")
	num, result = calcStr(num, result, 10, 50, 100, "X", "L", "C")
	num, result = calcStr(num, result, 1, 5, 10, "I", "V", "X")

	return result
}

func calc(i int, result string, str string) string {
	for i > 0 {
		result += str
		i--
	}
	return result
}

func calcStr(num int, result string, num1 int, num2 int, num3 int, str1 string, str2 string, str3 string) (int, string) {
	if num >= num1 && num < num2-num1 {
		i := num / num1
		result = calc(i, result, str1)
	} else if num >= num2-num1 && num < num2 {
		result += str1 + str2
	} else if num >= num2 && num < num3-num1 {
		result += str2
		i := num/num1 - 5
		result = calc(i, result, str1)
	} else if num >= num3-num1 {
		result += str1 + str3
	}
	return num % num1, result
}

// @lc code=end
