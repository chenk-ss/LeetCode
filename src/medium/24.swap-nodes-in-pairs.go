package medium

/*
 * @lc app=leetcode id=24 lang=golang
 *
 * [24] Swap Nodes in Pairs
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := &ListNode{-1, head}
	first := newHead.Next
	second := first.Next
	newHeadCp := newHead
	for first != nil && second != nil {
		last := second.Next
		first.Next = nil
		second.Next = first
		newHeadCp.Next = second
		newHeadCp = newHeadCp.Next.Next
		first = last
		if first == nil {
			return newHead.Next
		}
		second = first.Next
		if second == nil {
			newHeadCp.Next = first
			return newHead.Next
		}
	}
	return newHead.Next
}

// @lc code=end
