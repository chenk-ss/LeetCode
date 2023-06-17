package easy

import "strconv"

/*
 * @lc app=leetcode id=412 lang=golang
 *
 * [412] Fizz Buzz
 */

// @lc code=start
func fizzBuzz(n int) []string {
	result := make([]string, n)
	for n > 0 {
		if n%15 == 0 {
			result[n-1] = "FizzBuzz"
		} else if n%3 == 0 {
			result[n-1] = "Fizz"
		} else if n%5 == 0 {
			result[n-1] = "Buzz"
		} else {
			result[n-1] = strconv.Itoa(n)
		}
		n--
	}
	return result
}

// @lc code=end
