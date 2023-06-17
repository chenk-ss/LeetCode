package easy

/*
 * @lc app=leetcode id=232 lang=golang
 *
 * [232] Implement Queue using Stacks
 */

// @lc code=start
type MyQueue struct {
	list []int
	size int
}

func MyQueueConstructor() MyQueue {
	return *new(MyQueue)
}

func (this *MyQueue) Push(x int) {
	this.list = append(this.list, x)
	this.size++
}

func (this *MyQueue) Pop() int {
	x := this.list[0]
	this.list = this.list[1:]
	this.size--
	return x
}

func (this *MyQueue) Peek() int {
	return this.list[0]
}

func (this *MyQueue) Empty() bool {
	return this.size == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
// @lc code=end
