package medium

import "sort"

/*
 * @lc app=leetcode id=147 lang=golang
 *
 * [147] Insertion Sort List
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func insertionSortList(head *ListNode) *ListNode {
	nHead := &ListNode{}
	for head != nil {
		tmp := nHead
		for ; tmp.Next != nil && tmp.Next.Val < head.Val; tmp = tmp.Next {
		}
		tmp.Next, head.Next, head = head, tmp.Next, head.Next
	}
	return nHead.Next
}

func insertionSortList2(head *ListNode) *ListNode {
	list := []int{}
	for head != nil {
		list = append(list, head.Val)
		head = head.Next
	}
	sort.Ints(list)
	nHead := &ListNode{}
	node := nHead
	for _, v := range list {
		node.Next = &ListNode{Val: v}
		node = node.Next
	}
	return nHead.Next
}

// @lc code=end
