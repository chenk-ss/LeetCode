package easy

/*
 * @lc app=leetcode id=496 lang=golang
 *
 * [496] Next Greater Element I
 */

// @lc code=start
func nextGreaterElement(nums1 []int, nums2 []int) []int {
	result := []int{}
	for i := range nums1 {
		for j := range nums2 {
			if nums1[i] == nums2[j] {
				if j == len(nums2)-1 {
					result = append(result, -1)
				} else {
					flag := false
					for jj := j + 1; jj < len(nums2); jj++ {
						if nums1[i] < nums2[jj] {
							flag = true
							result = append(result, nums2[jj])
							break
						}
					}
					if !flag {
						result = append(result, -1)
					}
				}
			}
		}
	}
	return result
}

// @lc code=end
