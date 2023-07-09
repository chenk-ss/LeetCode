package medium

/*
 * @lc app=leetcode id=109 lang=golang
 *
 * [109] Convert Sorted List to Binary Search Tree
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedListToBST(head *ListNode) *TreeNode {
	list := []int{}
	for head != nil {
		list = append(list, head.Val)
		head = head.Next
	}
	return sortedListToBSTFunc(list)
}

func sortedListToBSTFunc(list []int) *TreeNode {
	if len(list) == 0 {
		return nil
	}
	mid := len(list) / 2
	root := &TreeNode{Val: list[mid]}
	if len(list) == 1 {
		return root
	}
	root.Left = sortedListToBSTFunc(list[:mid])
	root.Right = sortedListToBSTFunc(list[mid+1:])
	return root
}

// @lc code=end
