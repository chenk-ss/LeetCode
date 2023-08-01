package medium

/*
 * @lc app=leetcode id=229 lang=golang
 *
 * [229] Majority Element II
 */

// @lc code=start
func majorityElement(nums []int) []int {
	m := make(map[int]int)
	res := []int{}
	length := len(nums)
	for _, v := range nums {
		m[v]++
		if m[v] == length/3+1 {
			res = append(res, v)
		}
	}
	return res
}

// @lc code=end
