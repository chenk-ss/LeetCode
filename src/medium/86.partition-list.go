package medium

/*
 * @lc app=leetcode id=86 lang=golang
 *
 * [86] Partition List
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func partition(head *ListNode, x int) *ListNode {
	nHead, notBelowHead := &ListNode{}, &ListNode{}
	belowNode, notBelowNode := nHead, notBelowHead
	for head != nil {
		if head.Val < x {
			belowNode.Next = head
			belowNode = belowNode.Next
		} else {
			notBelowNode.Next = head
			notBelowNode = notBelowNode.Next
		}
		head = head.Next
	}
	notBelowNode.Next = nil
	belowNode.Next = notBelowHead.Next
	return nHead.Next
}

// @lc code=end
