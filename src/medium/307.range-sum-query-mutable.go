package medium

/*
 * @lc app=leetcode id=307 lang=golang
 *
 * [307] Range Sum Query - Mutable
 */

// @lc code=start
type NumArray struct {
	list []int
}

func Constructor(nums []int) NumArray {
	return NumArray{list: nums}
}

func (this *NumArray) Update(index int, val int) {
	this.list[index] = val
}

func (this *NumArray) SumRange(left int, right int) int {
	res := 0
	for i := left; i <= right; i++ {
		res += this.list[i]
	}
	return res
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * obj.Update(index,val);
 * param_2 := obj.SumRange(left,right);
 */
// @lc code=end
