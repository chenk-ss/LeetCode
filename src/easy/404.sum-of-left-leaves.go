package easy

/*
 * @lc app=leetcode id=404 lang=golang
 *
 * [404] Sum of Left Leaves
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
func sumOfLeftLeaves(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 0
	}
	return sumOfLeftLeavesFunc(root.Left, true) + sumOfLeftLeavesFunc(root.Right, false)
}

func sumOfLeftLeavesFunc(root *TreeNode, leftNode bool) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil && leftNode {
		return root.Val
	}
	return sumOfLeftLeavesFunc(root.Left, true) + sumOfLeftLeavesFunc(root.Right, false)
}

// @lc code=end
