package src

/*
 * @lc app=leetcode id=160 lang=golang
 *
 * [160] Intersection of Two Linked Lists
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	nodea, nodeb := headA, headB
	for nodea != nodeb {
		if nodea != nil {
			nodea = nodea.Next
		} else {
			nodea = headB
		}
		if nodeb != nil {
			nodeb = nodeb.Next
		} else {
			nodeb = headA
		}
	}
	return nodea
}

// @lc code=end
