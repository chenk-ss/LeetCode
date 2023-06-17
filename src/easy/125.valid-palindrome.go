package easy

import (
	"strings"
	"unicode"
)

/*
 * @lc app=leetcode id=125 lang=golang
 *
 * [125] Valid Palindrome
 */

// @lc code=start
func StringIsPalindrome(s string) bool {
	s = strings.ToLower(strings.ReplaceAll(s, " ", ""))
	left, right := 0, len(s)-1
	for left < right {
		leftChar, rightChar := s[left], s[right]
		if leftChar == rightChar {
			left++
			right--
		} else if !unicode.IsLetter(rune(leftChar)) && !unicode.IsNumber(rune(leftChar)) &&
			!unicode.IsLetter(rune(rightChar)) && !unicode.IsNumber(rune(rightChar)) {
			left++
			right--
		} else if !unicode.IsLetter(rune(leftChar)) && !unicode.IsNumber(rune(leftChar)) {
			left++
		} else if !unicode.IsLetter(rune(rightChar)) && !unicode.IsNumber(rune(rightChar)) {
			right--
		} else {
			return false
		}
	}
	return true
}

// @lc code=end
