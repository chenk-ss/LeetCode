package easy

/*
 * @lc app=leetcode id=349 lang=golang
 *
 * [349] Intersection of Two Arrays
 */

// @lc code=start
func intersection(nums1 []int, nums2 []int) []int {
	m := make(map[int]bool)
	for i := range nums1 {
		m[nums1[i]] = true
	}
	result := []int{}
	for i := range nums2 {
		if val, ok := m[nums2[i]]; ok && val {
			result = append(result, nums2[i])
			m[nums2[i]] = false
		}
	}
	return result
}

// @lc code=end
