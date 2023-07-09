package medium

/*
 * @lc app=leetcode id=97 lang=golang
 *
 * [97] Interleaving String
 */

// @lc code=start
func isInterleave(s1 string, s2 string, s3 string) bool {
	if len(s1)+len(s2) != len(s3) {
		return false
	}
	if s1 == "" || s2 == "" {
		return s1+s2 == s3
	}
	dp := make([][]bool, len(s1)+1)
	for i := range dp {
		dp[i] = make([]bool, len(s2)+1)
	}
	for i := 0; i < len(dp); i++ {
		for j := 0; j < len(dp[0]); j++ {
			if i == 0 || j == 0 {
				dp[i][j] = s1[:i]+s2[:j] == s3[:i+j]
				continue
			}
			if dp[i-1][j] && s1[i-1] == s3[i+j-1] {
				dp[i][j] = true
			}
			if dp[i][j-1] && s2[j-1] == s3[i+j-1] {
				dp[i][j] = true
			}
		}
	}
	return dp[len(s1)][len(s2)]
}

// @lc code=end
