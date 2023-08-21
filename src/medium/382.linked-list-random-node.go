package medium

import "math/rand"

/*
 * @lc app=leetcode id=382 lang=golang
 *
 * [382] Linked List Random Node
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type Solution struct {
	data []int
}

func Constructor(head *ListNode) Solution {
	data := []int{}
	for head != nil {
		data = append(data, head.Val)
		head = head.Next
	}
	return Solution{data}
}

func (this *Solution) GetRandom() int {
	return this.data[rand.Intn(len(this.data))]
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(head);
 * param_1 := obj.GetRandom();
 */
// @lc code=end
