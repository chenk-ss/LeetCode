package easy

import "strconv"

/*
 * @lc app=leetcode id=257 lang=golang
 *
 * [257] Binary Tree Paths
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
func binaryTreePaths(root *TreeNode) []string {
	return binaryTreePathsFunc(root, "")
}

func binaryTreePathsFunc(root *TreeNode, str string) []string {
	if root == nil {
		return nil
	}
	if root.Left == nil && root.Right == nil {
		return []string{str + strconv.Itoa(root.Val)}
	}
	str = str + strconv.Itoa(root.Val) + "->"
	return append(binaryTreePathsFunc(root.Left, str),
		binaryTreePathsFunc(root.Right, str)...)
}

// @lc code=end
