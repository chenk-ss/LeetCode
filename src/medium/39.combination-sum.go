package medium

/*
 * @lc app=leetcode id=39 lang=golang
 *
 * [39] Combination Sum
 */

// @lc code=start
func combinationSum(candidates []int, target int) [][]int {
	result := [][]int{}
	combinationSumFunc(candidates, target, []int{}, &result)
	return result
}

func combinationSumFunc(candidates []int, target int, former []int, result *[][]int) {
	if target == 0 {
		clone := make([]int, len(former))
		copy(clone, former)
		*result = append(*result, clone)
		return
	}
	if target > 0 && len(candidates) > 0 {
		combinationSumFunc(candidates[1:], target, former, result)
		v := candidates[0]
		combinationSumFunc(candidates, target-v, append(former, v), result)
	}
}

// @lc code=end
