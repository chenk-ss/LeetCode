package medium

/*
 * @lc app=leetcode id=122 lang=golang
 *
 * [117] Populating Next Right Pointers in Each Node II
 */

// @lc code=start

func maxProfit(prices []int) int {
	res := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			res += prices[i] - prices[i-1]
		}
	}
	return res
}

// @lc code=end
