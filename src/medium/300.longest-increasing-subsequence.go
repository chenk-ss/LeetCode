package medium

/*
 * @lc app=leetcode id=300 lang=golang
 *
 * [300] Longest Increasing Subsequence
 */

// @lc code=start
func lengthOfLIS(nums []int) int {
	dp := []int{}
	for _, v := range nums {
		if len(dp) == 0 || dp[len(dp)-1] < v {
			dp = append(dp, v)
		} else {
			index := lengthOfLISFunc(v, &dp)
			dp[index] = v
		}
	}
	return len(dp)
}

func lengthOfLISFunc(v int, dp *[]int) int {
	for i := range *dp {
		if (*dp)[i] >= v {
			return i
		}
	}
	return 0
}

// @lc code=end
