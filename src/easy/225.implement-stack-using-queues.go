package easy

/*
 * @lc app=leetcode id=225 lang=golang
 *
 * [225] Implement Stack using Queues
 */

// @lc code=start
type MyStack struct {
	list []int
	num  int
}

func constructor() MyStack {
	return *new(MyStack)
}

func (this *MyStack) Push(x int) {
	this.list = append(this.list, x)
	this.num++
}

func (this *MyStack) Pop() int {
	result := this.list[this.num-1]
	this.list = this.list[:this.num-1]
	this.num--
	return result
}

func (this *MyStack) Top() int {
	return this.list[this.num-1]
}

func (this *MyStack) Empty() bool {
	return this.num == 0
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
// @lc code=end
