package medium

import (
	"fmt"
	"math"
	"strings"
)

/*
 * @lc app=leetcode id=8 lang=golang
 *
 * [8] String to Integer (atoi)
 */

// @lc code=start
func myAtoi(s string) int {
	s = strings.TrimLeft(s, " ")

	if len(s) == 0 {
		return 0
	}
	result := 0
	negative := 1
	first := s[0]
	i := 0
	if first == '-' {
		negative = -1
		i++
	} else if first == '+' {
		i++
	}
	for i < len(s) {
		a := s[i]
		i++
		if a < '0' || a > '9' || result > math.MaxInt32 {
			break
		}
		result = result*10 + int(a) - '0'
	}
	result *= negative
	fmt.Println(result)
	if result > math.MaxInt32 {
		result = math.MaxInt32
	} else if result < math.MinInt32 {
		result = math.MinInt32
	}
	return result
}

// @lc code=end
