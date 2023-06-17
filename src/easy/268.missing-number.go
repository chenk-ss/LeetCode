package easy

/*
 * @lc app=leetcode id=268 lang=golang
 *
 * [268] Missing Number
 */

// @lc code=start
func missingNumber(nums []int) int {
	sum := 0
	for i := len(nums); i > 0; i-- {
		sum += i
	}
	for i := range nums {
		sum -= nums[i]
	}
	return sum
}

func missingNumber2(nums []int) int {
	list := make([]int, len(nums)+1)
	for i := range nums {
		list[nums[i]] = 1
	}
	for i := range list {
		if list[i] == 0 {
			return i
		}
	}
	return 0
}

// @lc code=end
