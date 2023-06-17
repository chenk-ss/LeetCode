package easy

/*
 * @lc app=leetcode id=217 lang=golang
 *
 * [217] Contains Duplicate
 */

// @lc code=start
func containsDuplicate(nums []int) bool {
	m := make(map[int]bool)
	for i := 0; i < len(nums); i++ {
		if _, ok := m[nums[i]]; ok {
			return true
		} else {
			m[nums[i]] = true
		}
	}
	return false
}

// @lc code=end
