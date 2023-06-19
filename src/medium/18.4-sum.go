package medium

import (
	"sort"
)

/*
 * @lc app=leetcode id=18 lang=golang
 *
 * [18] 4Sum
 */

// @lc code=start
func fourSum(nums []int, target int) [][]int {
	result := [][]int{}
	if len(nums) < 4 {
		return result
	}
	sort.Ints(nums)
	m := make(map[[4]int]bool)
	dd := [4]int{}
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			left, right := j+1, len(nums)-1
			for left < right {
				sum := nums[i] + nums[j] + nums[left] + nums[right]
				if sum == target {
					d := []int{nums[i], nums[j], nums[left], nums[right]}
					copy(dd[:], d)
					if _, ok := m[dd]; !ok {
						result = append(result, d)
						m[dd] = true
					}
					left++
					right--
					if left < right && nums[left-1] == nums[left] {
						left++
					}
					if left < right && nums[right+1] == nums[right] {
						right--
					}
				} else if sum > target {
					right--
				} else {
					left++
				}
			}
		}
	}

	return result
}

// @lc code=end
