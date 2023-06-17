package easy

/*
 * @lc app=leetcode id=500 lang=golang
 *
 * [500] Keyboard Row
 */

// @lc code=start
func findWords(words []string) []string {
	result := []string{}
	m := map[byte]int{
		'q': 0, 'w': 0, 'e': 0, 'r': 0, 't': 0, 'y': 0, 'u': 0, 'i': 0, 'o': 0, 'p': 0,
		'a': 1, 's': 1, 'd': 1, 'f': 1, 'g': 1, 'h': 1, 'j': 1, 'k': 1, 'l': 1,
		'z': 2, 'x': 2, 'c': 2, 'v': 2, 'b': 2, 'n': 2, 'm': 2,
	}
	for _, v := range words {
		flag := -1
		for _, b := range v {
			if b-'A' >= 0 && b-'A' < 26 {
				b = b - 'A' + 'a'
			}
			if flag == -1 {
				flag = m[byte(b)]
			}
			if flag != m[byte(b)] {
				flag = -1
				break
			}
		}
		if flag != -1 {
			result = append(result, v)
		}
	}
	return result
}

// @lc code=end
