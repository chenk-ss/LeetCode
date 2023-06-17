package easy

/*
 * @lc app=leetcode id=2405 lang=golang
 *
 * [2405] Optimal Partition of String
 */

// @lc code=start
func partitionString(s string) int {
	numList := make([]int, 26)
	count := 1
	for i := 0; i < len(s); i++ {
		j := s[i]
		if numList[j-'a'] == 1 {
			numList = make([]int, 26)
			count++
		}
		numList[j-'a'] = 1
	}
	return count
}

// @lc code=end
