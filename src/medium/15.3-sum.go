package medium

import "sort"

/*
 * @lc app=leetcode id=15 lang=golang
 *
 * [15] 3Sum
 */

// @lc code=start
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	result := [][]int{}
	numMap := map[int]int{}
	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		j := i + 1
		k := len(nums) - 1
		for j < k {
			sum := nums[i] + nums[j] + nums[k]
			if sum == 0 {
				data, contain := numMap[nums[i]]
				if (contain && data != nums[j]) || !contain {
					result = append(result, []int{nums[i], nums[j], nums[k]})
					numMap[nums[i]] = nums[j]
				}
				j++
				k--
			} else if sum > 0 {
				k--
			} else {
				j++
			}
		}
	}
	return result
}

// @lc code=end
