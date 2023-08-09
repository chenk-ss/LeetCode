package medium

/*
 * @lc app=leetcode id=442 lang=golang
 *
 * [442] Find All Duplicates in an Array
 */

// @lc code=start
func findDuplicates(nums []int) []int {
	list := make([]int, len(nums)+1)
	for _, v := range nums {
		list[v]++
	}
	res := []int{}
	for i, v := range list {
		if v > 1 {
			res = append(res, i)
		}
	}
	return res
}

// @lc code=end
