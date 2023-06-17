package easy

/*
 * @lc app=leetcode id=202 lang=golang
 *
 * [202] Happy Number
 */

// @lc code=start
func isHappy(n int) bool {

	nums := []int{}
	for n != 1 {
		for i := 0; i < len(nums)-1; i++ {
			if nums[i] == n {
				return false
			}
		}
		sum := 0
		for n > 0 {
			sum += (n % 10) * (n % 10)
			n /= 10
		}
		n = sum
		nums = append(nums, sum)
	}
	return true
}

// @lc code=end
