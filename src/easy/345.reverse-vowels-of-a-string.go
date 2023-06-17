package easy

/*
 * @lc app=leetcode id=345 lang=golang
 *
 * [345] Reverse Vowels of a String
 */

// @lc code=start
func reverseVowels(s string) string {
	left, right := 0, len(s)-1
	m := map[byte]int{
		'a': 1,
		'e': 1,
		'i': 1,
		'o': 1,
		'u': 1,
		'A': 1,
		'E': 1,
		'I': 1,
		'O': 1,
		'U': 1,
	}
	resultList := make([]byte, len(s))
	for left <= right {
		leftByte, rightByte := s[left], s[right]
		_, ok1 := m[leftByte]
		_, ok2 := m[rightByte]
		if !ok1 {
			resultList[left] = leftByte
			left++
		}
		if !ok2 {
			resultList[right] = rightByte
			right--
		}
		if ok1 && ok2 {
			resultList[left], resultList[right] = rightByte, leftByte
			left++
			right--
		}
	}
	return string(resultList)
}

// @lc code=end
