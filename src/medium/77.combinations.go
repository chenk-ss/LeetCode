package medium

/*
 * @lc app=leetcode id=77 lang=golang
 *
 * [77] Combinations
 */

// @lc code=start
func combine(n int, k int) [][]int {
	return combineFunc(1, n, k)
}

func combineFunc(start, n int, k int) [][]int {
	result := [][]int{}
	if k == 1 {
		for i := start; i <= n; i++ {
			result = append(result, []int{i})
		}
		return result
	}
	for start <= n-k+1 {
		data := combineFunc(start+1, n, k-1)
		for _, v := range data {
			result = append(result, append([]int{start}, v...))
		}
		start++
	}
	return result
}

// @lc code=end
