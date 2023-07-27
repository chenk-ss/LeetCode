package medium

/*
 * @lc app=leetcode id=204 lang=golang
 *
 * [204] Count Primes
 */

// @lc code=start
func countPrimes(n int) int {
	if n < 3 {
		return 0
	}
	arr := make([]bool, n+1)
	for i := 2; i < n; i++ {
		if !arr[i] {
			for j := 2; i*j <= n; j++ {
				arr[i*j] = true
			}
		}
	}
	res := 0
	for i, v := range arr {
		if !v && i > 1 && i != n {
			res++
		}
	}
	return res
}

// @lc code=end
