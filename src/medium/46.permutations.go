package medium

/*
 * @lc app=leetcode id=46 lang=golang
 *
 * [46] Permutations
 */

// @lc code=start
func permute(nums []int) [][]int {
	result := [][]int{}
	if len(nums) == 1 {
		result = append(result, nums)
		return result
	}
	for i := range nums {
		param := append(append([]int{}, nums[:i]...), nums[i+1:]...)
		for _, v := range permute(param) {
			result = append(result, append([]int{nums[i]}, v...))
		}
	}
	return result
}

// @lc code=end
