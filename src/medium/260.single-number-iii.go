package medium

/*
 * @lc app=leetcode id=260 lang=golang
 *
 * [260] Single Number III
 */

// @lc code=start
func singleNumber(nums []int) []int {
	if len(nums) == 2 {
		return nums
	}
	m := make(map[int]bool)
	res := []int{}
	for _, v := range nums {
		if _, ok := m[v]; ok {
			delete(m, v)
		} else {
			m[v] = true
		}
	}
	for k := range m {
		res = append(res, k)
	}
	return res
}

// @lc code=end
