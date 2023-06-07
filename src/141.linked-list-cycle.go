package src

/*
 * @lc app=leetcode id=141 lang=golang
 *
 * [141] Linked List Cycle
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	oneStepNode, twoStepNode := head, head.Next
	for oneStepNode != nil && twoStepNode != nil {
		if oneStepNode == twoStepNode {
			return true
		}
		if twoStepNode.Next == nil {
			return false
		}
		oneStepNode = oneStepNode.Next
		twoStepNode = twoStepNode.Next.Next
	}
	return false
}

// @lc code=end
