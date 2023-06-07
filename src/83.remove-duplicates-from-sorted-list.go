package src

/*
 * @lc app=leetcode id=83 lang=golang
 *
 * [83] Remove Duplicates from Sorted List
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	former := head
	next := former.Next
	for next != nil {
		if former.Val == next.Val {
			next = next.Next
		} else {
			former.Next = next
			former = next
			next = next.Next
		}
	}
	former.Next = nil
	return head
}

// @lc code=end
