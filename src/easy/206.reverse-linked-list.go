package easy

/*
 * @lc app=leetcode id=206 lang=golang
 *
 * [206] Reverse Linked List
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	var result *ListNode
	for head != nil {
		second := head.Next
		head.Next = result
		result = head
		head = second
	}
	return result
}

// @lc code=end
