package easy

import "math"

/*
 * @lc app=leetcode id=530 lang=golang
 *
 * [530] Minimum Absolute Difference in BST
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
func getMinimumDifference(root *TreeNode) int {
	min := math.MaxInt
	var pre *TreeNode
	var getMinimumDifferenceFunc func(root *TreeNode)
	getMinimumDifferenceFunc = func(root *TreeNode) {
		if root == nil {
			return
		}
		getMinimumDifferenceFunc(root.Left)
		if pre != nil && min > root.Val-pre.Val {
			min = root.Val - pre.Val
		}
		pre = root
		getMinimumDifferenceFunc(root.Right)
	}
	getMinimumDifferenceFunc(root)
	return min
}

// @lc code=end
