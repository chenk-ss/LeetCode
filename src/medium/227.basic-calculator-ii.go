package medium

/*
 * @lc app=leetcode id=227 lang=golang
 *
 * [227] Basic Calculator II
 */

// @lc code=start
func calculate(s string) int {
	list := []int{}
	num := 0
	op := '+'
	s += "+"
	for _, v := range s {
		if v == ' ' {
			continue
		} else if v >= '0' && v <= '9' {
			num = num*10 + int(v-'0')
		} else {
			switch op {
			case '+':
				list = append(list, num)
			case '-':
				list = append(list, -num)
			case '*':
				list[len(list)-1] *= num
			case '/':
				list[len(list)-1] /= num
			}
			num = 0
			op = v
		}
	}
	res := 0
	for _, v := range list {
		res += v
	}
	return res
}

// @lc code=end
