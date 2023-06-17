package easy

/*
 * @lc app=leetcode id=441 lang=golang
 *
 * [441] Arranging Coins
 */

// @lc code=start
func arrangeCoins(n int) int {
	left, mid, right := 1, 1, n
	for left <= right {
		mid = (left + right) / 2
		sum := (mid + 1) * mid / 2
		if sum > n {
			right = mid - 1
		} else if sum < n {
			left = mid + 1
		} else {
			return mid
		}
	}
	return right
}

// @lc code=end
