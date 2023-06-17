package easy

/*
 * @lc app=leetcode id=219 lang=golang
 *
 * [219] Contains Duplicate II
 */

// @lc code=start
func containsNearbyDuplicate(nums []int, k int) bool {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if j, ok := m[nums[i]]; ok {
			if i-j <= k {
				return true
			}
		}
		m[nums[i]] = i
	}
	return false
}

// @lc code=end
