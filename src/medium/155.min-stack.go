package medium

import "math"

/*
 * @lc app=leetcode id=155 lang=golang
 *
 * [155] Min Stack
 */

// @lc code=start
type MinStack struct {
	min  int
	vals []int
}

func Constructor155() MinStack {
	return MinStack{min: math.MaxInt}
}

func (this *MinStack) Push(val int) {
	this.vals = append(this.vals, val)
	if val < this.min {
		this.min = val
	}
}

func (this *MinStack) Pop() {
	if this.vals[len(this.vals)-1] == this.min {
		this.min = math.MaxInt
		for i := 0; i < len(this.vals)-1; i++ {
			if this.vals[i] < this.min {
				this.min = this.vals[i]
			}
		}
	}
	this.vals = this.vals[:len(this.vals)-1]
}

func (this *MinStack) Top() int {
	return this.vals[len(this.vals)-1]
}

func (this *MinStack) GetMin() int {
	return this.min
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
// @lc code=end
