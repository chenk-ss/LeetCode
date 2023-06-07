package src

/*
 * @lc app=leetcode id=70 lang=golang
 *
 * [70] Climbing Stairs
 */

// @lc code=start
func climbStairs(n int) int {
	if n == 1 || n == 2 {
		return n
	}
	list := make([]int, n)
	list[0] = 1
	list[1] = 2
	for i := 2; i < n; i++ {
		list[i] = list[i-1] + list[i-2]
	}
	return list[n-1]
}

// @lc code=end
