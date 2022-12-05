package src

/*
 * @lc app=leetcode id=13 lang=golang
 *
 * [13] Roman to Integer
 */

// @lc code=start
func romanToInt(s string) int {
	m := map[byte]int{
		'M': 1000,
		'D': 500,
		'C': 100,
		'L': 50,
		'X': 10,
		'V': 5,
		'I': 1,
	}
	result := 0
	i := 0
	for i < len(s) {
		str := s[i]
		num := m[str]
		if i+1 < len(s) {
			str1 := s[i+1]
			num1 := m[str1]
			if num < num1 {
				result += num1 - num
				i += 2
			} else {
				result += num
				i++
			}
		} else {
			result += num
			i++
		}
	}
	return result
}

// @lc code=end
