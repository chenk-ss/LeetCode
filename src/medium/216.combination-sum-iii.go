package medium

/*
 * @lc app=leetcode id=216 lang=golang
 *
 * [216] Combination Sum III
 */

// @lc code=start
func combinationSum3(k int, n int) [][]int {
	res := [][]int{}
	combinationSum3Func(k, n, 0, &res, []int{})
	return res
}

func combinationSum3Func(k, n, max int, res *[][]int, list []int) {
	if len(list) == k && n == 0 {
		*res = append(*res, append([]int{}, list...))
		return
	}
	if len(list) == k || n <= 0 {
		return
	}
	for i := max + 1; i < 10; i++ {
		combinationSum3Func(k, n-i, i, res, append(list, i))
	}
}

// @lc code=end
