package medium

import "sort"

/*
 * @lc app=leetcode id=31 lang=golang
 *
 * [31] Next Permutation
 */

// @lc code=start
func nextPermutation(nums []int) {
	if len(nums) == 1 {
		return
	}
	k := 0
	kv := 101
	for i := len(nums) - 2; i >= 0; i-- {
		iv := nums[i]
		for j := i + 1; j < len(nums); j++ {
			jv := nums[j]
			if iv < jv && jv < kv {
				k = j
				kv = jv
			}
		}
		if k != 0 {
			nums[i], nums[k] = nums[k], nums[i]
			bb := nums[i+1:]
			sort.Ints(bb)
			nums = append(nums[:i+1], bb...)
			return
		}
	}
	if k == 0 {
		sort.Ints(nums)
	}
}

// @lc code=end
