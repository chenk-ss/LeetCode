package medium

/*
 * @lc app=leetcode id=264 lang=golang
 *
 * [264] Ugly Number II
 */

// @lc code=start
func nthUglyNumber(n int) int {
	res := make([]int, n)
	res[0] = 1
	i2, i3, i5 := 0, 0, 0
	for i := 1; i < n; i++ {
		min := nthUglyNumberFunc(res[i2]*2, nthUglyNumberFunc(res[i3]*3, res[i5]*5))
		res[i] = min
		if min == res[i2]*2 {
			i2++
		}
		if min == res[i3]*3 {
			i3++
		}
		if min == res[i5]*5 {
			i5++
		}
	}
	return res[n-1]
}

func nthUglyNumberFunc(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// @lc code=end
