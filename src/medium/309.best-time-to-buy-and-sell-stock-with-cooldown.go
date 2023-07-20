package medium

/*
 * @lc app=leetcode id=309 lang=golang
 *
 * [309] Best Time to Buy and Sell Stock with Cooldown
 */

// @lc code=start
func maxProfit(prices []int) int {
	sell, buy, rest := 0, -prices[0], 0
	for i := 1; i < len(prices); i++ {
		sell, buy, rest = maxProfitFunc(sell, buy+prices[i]), maxProfitFunc(buy, rest-prices[i]), maxProfitFunc(rest, sell)
	}
	return maxProfitFunc(sell, maxProfitFunc(buy, rest))
}

func maxProfitFunc(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
