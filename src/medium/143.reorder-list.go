package medium

/*
 * @lc app=leetcode id=143 lang=golang
 *
 * [143] Reorder List
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reorderList(head *ListNode) {
	if head == nil {
		return
	}
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	last := slow.Next
	slow.Next = nil
	var lastReverse *ListNode
	for last != nil {
		last.Next, lastReverse, last = lastReverse, last, last.Next
	}
	node1, node2 := head, lastReverse
	for node1 != nil && node2 != nil {
		next1, next2 := node1.Next, node2.Next
		node1.Next = node2
		node1.Next.Next = next1
		node1, node2 = next1, next2
	}
}

// @lc code=end
