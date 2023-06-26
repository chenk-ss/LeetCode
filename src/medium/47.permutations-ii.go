package medium

import "sort"

/*
 * @lc app=leetcode id=47 lang=golang
 *
 * [47] Permutations II
 */

// @lc code=start
func permuteUnique(nums []int) [][]int {
	sort.Ints(nums)
	return permuteUniqueFunc(nums)
}

func permuteUniqueFunc(nums []int) [][]int {
	result := [][]int{}
	if len(nums) == 1 {
		result = append(result, nums)
		return result
	}
	for i := range nums {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		param := append(append([]int{}, nums[:i]...), nums[i+1:]...)
		for _, v := range permuteUniqueFunc(param) {
			result = append(result, append([]int{nums[i]}, v...))
		}
	}
	return result
}

// @lc code=end
