package main

/*
 * @lc app=leetcode id=1 lang=golang
 *
 * [1] Two Sum
 */

// @lc code=start
func twoSum(nums []int, target int) []int {
	if len(nums) == 0 {
		return nil
	}
	result := make([]int, 2)
	myMap := make(map[int]int, len(nums))
	for i := 0; i < len(nums); i++ {
		if val, ok := myMap[target-nums[i]]; ok {
			result[0] = val
			result[1] = i
			return result
		} else {
			myMap[nums[i]] = i
		}
	}
	return result
}

// @lc code=end
