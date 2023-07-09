package medium

/*
 * @lc app=leetcode id=96 lang=golang
 *
 * [96] Unique Binary Search Trees
 */

// @lc code=start
func numTrees(n int) int {
	list := make([]int, n+1)
	list[0], list[1] = 1, 1
	for i := 2; i <= n; i++ {
		for j := 0; j < i; j++ {
			list[i] += list[j] * list[i-j-1]
		}
	}
	return list[n]
}

// @lc code=end
