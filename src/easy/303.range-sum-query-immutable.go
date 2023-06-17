package easy

/*
 * @lc app=leetcode id=303 lang=golang
 *
 * [303] Range Sum Query - Immutable
 */

// @lc code=start
type NumArray struct {
	list []int
}

func Constructor(nums []int) NumArray {
	array := new(NumArray)
	array.list = nums
	return *array
}

func (this *NumArray) SumRange(left int, right int) int {
	result := 0
	for i, list := 0, this.list[left:right+1]; i < len(list); i++ {
		result += list[i]
	}
	return result
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */
// @lc code=end
