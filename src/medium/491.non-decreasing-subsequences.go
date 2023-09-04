package medium

/*
 * @lc app=leetcode id=491 lang=golang
 *
 * [491] Non-decreasing Subsequences
 */

// @lc code=start
func findSubsequences(nums []int) [][]int {
	res := [][]int{}
	findSubsequencesFunc(0, nums, []int{}, &res)
	return res
}

func findSubsequencesFunc(index int, nums []int, list []int, res *[][]int) {
	if len(list) > 1 {
		*res = append(*res, append([]int{}, list...))
	}
	m := make(map[int]bool)
	for i := index; i < len(nums); i++ {
		if m[nums[i]] || (len(list) > 0 && (list)[len(list)-1] > nums[i]) {
			continue
		}
		m[nums[i]] = true
		list = append(list, nums[i])
		findSubsequencesFunc(i+1, nums, list, res)
		list = (list)[:len(list)-1]
	}
}

// @lc code=end
