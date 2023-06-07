package src

import (
	"math"
	"sort"
)

/*
 * @lc app=leetcode id=16 lang=golang
 *
 * [16] 3Sum Closest
 */

// @lc code=start
func threeSumClosest(nums []int, target int) int {
	closest := nums[0] + nums[1] + nums[2]
	diff := math.Abs(float64(closest) - float64(target))
	sort.Ints(nums)
	for i := 0; i < len(nums)-2; i++ {
		left := i + 1
		right := len(nums) - 1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			newDiff := math.Abs(float64(sum) - float64(target))
			if diff > newDiff {
				diff = newDiff
				closest = sum
			}
			if sum < target {
				left++
			}
		}
	}
	return closest
}

// @lc code=end
