package medium

/*
 * @lc app=leetcode id=114 lang=golang
 *
 * [114] Flatten Binary Tree to Linked List
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	temp := root.Right
	root.Right = root.Left
	root.Left = nil

	right := root
	for right.Right != nil {
		right = right.Right
	}
	right.Right = temp
	flatten(root.Left)
	flatten(root.Right)
}

// @lc code=end
