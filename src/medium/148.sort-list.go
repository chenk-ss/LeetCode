package medium

import "sort"

/*
 * @lc app=leetcode id=148 lang=golang
 *
 * [148] Sort List
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func sortList(head *ListNode) *ListNode {
	list := []int{}
	for head != nil {
		list = append(list, head.Val)
		head = head.Next
	}
	sort.Ints(list)
	res := &ListNode{}
	node := res
	for _, v := range list {
		node.Next = &ListNode{Val: v}
		node = node.Next
	}
	return res.Next
}

// @lc code=end
