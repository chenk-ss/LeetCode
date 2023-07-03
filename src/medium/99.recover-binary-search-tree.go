package medium

/*
 * @lc app=leetcode id=99 lang=golang
 *
 * [99] Recover Binary Search Tree
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
func recoverTree(root *TreeNode) {
	var pre, first, second *TreeNode
	var recoverTreeFunc func(node *TreeNode)
	recoverTreeFunc = func(node *TreeNode) {
		if node == nil {
			return
		}
		recoverTreeFunc(node.Left)
		if pre != nil && pre.Val >= node.Val {
			if first == nil {
				first = pre
			}
			if first != nil {
				second = node
			}
		}
		pre = node
		recoverTreeFunc(node.Right)
	}
	recoverTreeFunc(root)
	first.Val, second.Val = second.Val, first.Val
}

// @lc code=end
