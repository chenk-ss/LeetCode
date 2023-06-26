package medium /*
 * @lc app=leetcode id=40 lang=golang
 *
 * [40] Combination Sum II
 */

import "sort"

// @lc code=start
func combinationSum2(candidates []int, target int) [][]int {
	result := [][]int{}
	sort.Ints(candidates)
	combinationSum2Func(candidates, target, []int{}, &result)
	return result
}

func combinationSum2Func(candidates []int, target int, former []int, result *[][]int) {
	if target == 0 {
		*result = append(*result, append([]int{}, former...))
		return
	}
	for i := 0; i < len(candidates); i++ {
		if candidates[i] > target {
			break
		}
		if i > 0 && candidates[i] == candidates[i-1] {
			continue
		}
		combinationSum2Func(candidates[i+1:], target-candidates[i], append(former, candidates[i]), result)
	}
}

// @lc code=end
