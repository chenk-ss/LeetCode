package medium

/*
 * @lc app=leetcode id=128 lang=golang
 *
 * [128] Longest Consecutive Sequence
 */

// @lc code=start
func longestConsecutive(nums []int) int {
	maxLen := 0
	m := make(map[int]bool)
	for _, v := range nums {
		m[v] = true
	}
	for num := range m {
		if m[num-1] {
			continue
		}
		cur := num
		for m[cur+1] {
			cur++
		}
		if cur-num+1 > maxLen {
			maxLen = cur - num + 1
		}
	}
	return maxLen
}

// @lc code=end
