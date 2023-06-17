package easy

/*
 * @lc app=leetcode id=350 lang=golang
 *
 * [350] Intersection of Two Arrays II
 */

// @lc code=start
func intersect(nums1 []int, nums2 []int) []int {
	m := make(map[int]int)
	for i := range nums1 {
		if val, ok := m[nums1[i]]; ok {
			m[nums1[i]] = val + 1
		} else {
			m[nums1[i]] = 1
		}
	}
	result := []int{}
	for i := range nums2 {
		if val, ok := m[nums2[i]]; ok && val > 0 {
			result = append(result, nums2[i])
			m[nums2[i]] = val - 1
		}
	}
	return result
}

// @lc code=end
