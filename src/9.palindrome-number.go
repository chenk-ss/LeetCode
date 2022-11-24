package src

/*
 * @lc app=leetcode id=9 lang=golang
 *
 * [9] Palindrome Number
 */

// @lc code=start
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	originX := x
	rev := 0
	last := 0
	for originX > 0 {
		last = originX % 10
		rev = rev*10 + last
		originX /= 10
	}
	return x == rev
}

// @lc code=end
