package easy

/*
 * @lc app=leetcode id=121 lang=golang
 *
 * [121] Best Time to Buy and Sell Stock
 */

// @lc code=start
func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	left, max := prices[0], 0
	for i := 0; i < len(prices); i++ {
		if left > prices[i] {
			left = prices[i]
		}
		if prices[i]-left > max {
			max = prices[i] - left
		}
	}
	return max
}

// @lc code=end
