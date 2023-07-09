package medium

import (
	"math"
	"sort"
)

/*
 * @lc app=leetcode id=373 lang=golang
 *
 * [373] Find K Pairs with Smallest Sums
 */

// @lc code=start
func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	res := [][]int{}
	m, n := len(nums1)-1, len(nums2)-1
	maxValue := nums1[m] + nums2[n] + 1
	// milestone - index of value that never will be in answers
	// For example if k = 4
	// nums1[2]+nums2[2] will never exist in answer
	milestone := int(math.Sqrt(float64(k))) + 1
	if milestone < m && milestone < n {
		maxValue = nums1[milestone] + nums2[milestone]
	}

	for i := 0; i < len(nums1); i++ {
		for j := 0; j < len(nums2); j++ {
			if nums1[i]+nums2[j] > maxValue {
				break
			}
			res = append(res, []int{nums1[i], nums2[j]})
		}
	}
	sort.Slice(res, func(i, j int) bool {
		return res[i][0]+res[i][1] < res[j][0]+res[j][1]
	})
	if len(res) > k {
		res = res[:k]
	}
	return res
}

// @lc code=end
