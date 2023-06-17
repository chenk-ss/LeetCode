package easy

import (
	"strconv"
)

/*
 * @lc app=leetcode id=228 lang=golang
 *
 * [228] Summary Ranges
 */

// @lc code=start
func summaryRanges(nums []int) []string {
	result := []string{}
	for i, j := 0, 0; i < len(nums); {
		a, b := nums[i], nums[j]
		if j < len(nums)-1 {
			j++
		}
		step := 1
		for j < len(nums) && nums[j] == nums[i]+step {
			b = nums[j]
			j++
			step++
		}
		if a == b {
			result = append(result, strconv.Itoa(a))
			i++
		} else {
			result = append(result, strconv.Itoa(a)+"->"+strconv.Itoa(b))
			i = j
		}
	}
	return result
}

// @lc code=end
