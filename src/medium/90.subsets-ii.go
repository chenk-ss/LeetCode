package medium

import "sort"

/*
 * @lc app=leetcode id=90 lang=golang
 *
 * [90] Subsets II
 */

// @lc code=start
func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	result := [][]int{}
	for i := 0; i <= len(nums); i++ {
		subsetsWithDupFun(nums, i, 0, []int{}, &result)
	}
	return result
}

func subsetsWithDupFun(nums []int, k int, index int, list []int, result *[][]int) {
	if len(list) == k {
		*result = append(*result, append([]int{}, list...))
	}
	for i := index; i < len(nums); i++ {
		if i != index && nums[i] == nums[i-1] {
			continue
		}
		subsetsWithDupFun(nums, k, i+1, append(list, nums[i]), result)
	}
}

// @lc code=end
