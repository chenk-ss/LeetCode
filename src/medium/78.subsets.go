package medium

/*
 * @lc app=leetcode id=78 lang=golang
 *
 * [78] Subsets
 */

// @lc code=start
func subsets(nums []int) [][]int {
	result := [][]int{}
	result = append(result, []int{})
	for i := 1; i <= len(nums); i++ {
		data := subsetsFun(nums, i)
		for _, v := range data {
			result = append(result, v)
		}
	}
	return result
}

func subsetsFun(nums []int, k int) [][]int {
	result := [][]int{}
	if k == 1 {
		for _, v := range nums {
			result = append(result, []int{v})
		}
		return result
	}
	for len(nums) >= k {
		data := subsetsFun(nums[1:], k-1)
		for _, v := range data {
			result = append(result, append([]int{nums[0]}, v...))
		}
		nums = nums[1:]
	}
	return result
}

// @lc code=end
