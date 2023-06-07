package src

/*
 * @lc app=leetcode id=110 lang=golang
 *
 * [110] Balanced Binary Tree
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
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	leftDepth, rightDepth := isBalancedMaxDepth(root.Left), isBalancedMaxDepth(root.Right)
	if isBalancedAbs(leftDepth, rightDepth) > 1 {
		return false
	}
	return isBalanced(root.Left) && isBalanced(root.Right)
}

func isBalancedMaxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftDepth, rightDepth := isBalancedMaxDepth(root.Left), isBalancedMaxDepth(root.Right)
	return 1 + isBalancedMax(leftDepth, rightDepth)
}

func isBalancedAbs(a int, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

func isBalancedMax(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
