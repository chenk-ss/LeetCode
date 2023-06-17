package easy

/*
 * @lc app=leetcode id=4 lang=golang
 *
 * [4] Median of Two Sorted Arrays
 */

// @lc code=start
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m := len(nums1)
	n := len(nums2)
	i := 0
	j := 0
	tmp := (m + n) % 2
	mid := (m + n + 1) / 2
	num1 := 0
	num2 := 0
	cnt := 0
	for {
		if cnt == mid+1-tmp {
			break
		}
		cnt++
		num1 = num2
		if i == m {
			num2 = nums2[j]
			j++
			continue
		}
		if j == n {
			num2 = nums1[i]
			i++
			continue
		}
		if nums1[i] >= nums2[j] {
			num2 = nums2[j]
			j++
		} else {
			num2 = nums1[i]
			i++
		}
	}
	if tmp == 0 {
		return float64(num1+num2) / 2
	} else {
		return float64(num2)
	}
}

// @lc code=end
