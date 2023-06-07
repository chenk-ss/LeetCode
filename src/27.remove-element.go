package src

/*
 * @lc app=leetcode id=27 lang=golang
 *
 * [27] Remove Element
 */

// @lc code=start
func removeElement(nums []int, val int) int {
	k := 0
	for i := 0; i < len(nums); i++ {
		element := nums[i]
		if element != val {
			nums[k] = nums[i]
			k++
		}
	}
	return k
}

// @lc code=end
