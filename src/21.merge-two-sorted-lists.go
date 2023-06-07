package src

/*
 * @lc app=leetcode id=21 lang=golang
 *
 * [21] Merge Two Sorted Lists
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	result := &ListNode{}
	data := result
	for list1 != nil && list2 != nil {
		if list1.Val > list2.Val {
			data.Next = list2
			list2 = list2.Next
		} else {
			data.Next = list1
			list1 = list1.Next
		}
		data = data.Next
	}
	if list1 == nil {
		data.Next = list2
	}
	if list2 == nil {
		data.Next = list1
	}
	return result.Next
}

// @lc code=end
