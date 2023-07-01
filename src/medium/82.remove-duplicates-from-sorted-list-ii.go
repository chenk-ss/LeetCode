package medium

/*
 * @lc app=leetcode id=82 lang=golang
 *
 * [82] Remove Duplicates from Sorted List II
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
	if head == nil || head.Next == nil {
		return head
	}
	nHead := &ListNode{}
	node := nHead
	for head != nil {
		val := head.Val
		count := 1
		for head.Next != nil && val == head.Next.Val {
			head = head.Next
			count++
		}
		if count == 1 {
			node.Next = head
			node = node.Next
		}
		head = head.Next
	}
	node.Next = nil
	return nHead.Next
}

// @lc code=end
