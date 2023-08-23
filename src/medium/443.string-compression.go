package medium

import "strconv"

/*
 * @lc app=leetcode id=443 lang=golang
 *
 * [443] String Compression
 */

// @lc code=start
func compress(chars []byte) int {
	count, index := 1, 0
	for i := 1; i <= len(chars); i++ {
		if i < len(chars) && chars[i] == chars[i-1] {
			count++
		} else {
			chars[index] = chars[i-1]
			index++
			if count > 1 {
				cnt := strconv.Itoa(count)
				for _, v := range cnt {
					chars[index] = byte(v)
					index++
				}
			}
			count = 1
		}
	}
	return index
}

// @lc code=end
