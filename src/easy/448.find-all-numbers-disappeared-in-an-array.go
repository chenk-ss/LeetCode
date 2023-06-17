package easy

/*
 * @lc app=leetcode id=448 lang=golang
 *
 * [448] Find All Numbers Disappeared in an Array
 */

// @lc code=start
func findDisappearedNumbers(nums []int) []int {
	list := make([]int, len(nums))
	for i := range nums {
		list[nums[i]-1] = 1
	}
	result := []int{}
	for i := range list {
		if list[i] == 0 {
			result = append(result, i+1)
		}
	}
	return result
}

// @lc code=end
