package medium

/*
 * @lc app=leetcode id=45 lang=golang
 *
 * [45] Jump Game II
 */

// @lc code=start
func jump(nums []int) int {
	if len(nums) < 2 {
		return 0
	}
	maxIdx := len(nums) - 1
	list := make([]int, len(nums))
	for i := range list {
		list[i] = len(nums)
	}
	for i := maxIdx; i >= 0; i-- {
		if i+nums[i] >= maxIdx {
			list[i] = 1
		} else {
			for j := i + 1; j <= i+nums[i]; j++ {
				if list[i] > list[j]+1 {
					list[i] = list[j] + 1
				}
			}
		}
	}
	return list[0]
}

func jump2(nums []int) int {
	count, cur, maxJump := 0, 0, 0
	for i := 0; i < len(nums)-1; i++ {
		if i+nums[i] > maxJump {
			maxJump = i + nums[i]
		}
		if i == cur {
			cur, count = maxJump, count+1
			if cur >= len(nums)-1 {
				return count
			}
		}
	}
	return 0
}

// @lc code=end
