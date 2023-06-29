package medium

/*
 * @lc app=leetcode id=75 lang=golang
 *
 * [75] Sort Colors
 */

// @lc code=start
func sortColors(nums []int) {
	left, right := 0, len(nums)-1
	index := 0
	for index <= right {
		if nums[index] == 0 {
			nums[index], nums[left] = nums[left], nums[index]
			left++
			index++
		} else if nums[index] == 2 {
			nums[index], nums[right] = nums[right], nums[index]
			right--
		} else {
			index++
		}
	}
}

func sortColors2(nums []int) {
	count0, count1 := 0, 0
	for _, v := range nums {
		if v == 0 {
			count0++
		} else if v == 1 {
			count1++
		}
	}
	for i := range nums {
		if count0 > 0 {
			nums[i] = 0
			count0--
			continue
		}
		if count1 > 0 {
			nums[i] = 1
			count1--
			continue
		}
		nums[i] = 2
	}
}

// @lc code=end
