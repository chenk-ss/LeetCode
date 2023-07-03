package medium

import "math"

/*
 * @lc app=leetcode id=98 lang=golang
 *
 * [98] Validate Binary Search Tree
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
func isValidBST(root *TreeNode) bool {
	return isValidBSTFunc(root, math.MinInt, math.MaxInt)
}

func isValidBSTFunc(node *TreeNode, min, max int) bool {
	if node == nil {
		return true
	}
	if node.Val > min && node.Val < max {
		return isValidBSTFunc(node.Left, min, node.Val) && isValidBSTFunc(node.Right, node.Val, max)
	} else {
		return false
	}
}

// @lc code=end
