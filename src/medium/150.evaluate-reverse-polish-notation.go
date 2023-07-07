package medium

import "strconv"

/*
 * @lc app=leetcode id=150 lang=golang
 *
 * [150] Evaluate Reverse Polish Notation
 */

// @lc code=start
func evalRPN(tokens []string) int {
	list := make([]int, len(tokens))
	length := 0
	for _, v := range tokens {
		if v != "+" && v != "-" && v != "*" && v != "/" {
			if val, e := strconv.Atoi(v); e == nil {
				list[length] = val
				length++
			}
			continue
		}
		first, second := list[length-2], list[length-1]
		next := 0
		if v == "+" {
			next = first + second
		} else if v == "-" {
			next = first - second
		} else if v == "*" {
			next = first * second
		} else if v == "/" {
			next = first / second
		}
		list[length-2] = next
		length--
	}
	return list[0]
}

// @lc code=end
