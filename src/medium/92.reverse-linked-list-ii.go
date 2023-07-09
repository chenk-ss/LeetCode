package medium

/*
 * @lc app=leetcode id=92 lang=golang
 *
 * [92] Reverse Linked List II
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if head == nil {
		return head
	}
	nHead := &ListNode{Next: head}
	leftLine, cur := nHead, nHead
	count := 0
	for ; count < left; count++ {
		leftLine, cur = cur, cur.Next
	}
	var reverseLine *ListNode
	for ; count <= right; count++ {
		reverseLine, cur, cur.Next = cur, cur.Next, reverseLine
	}
	leftLine.Next, leftLine.Next.Next = reverseLine, cur
	return nHead.Next
}

// @lc code=end
