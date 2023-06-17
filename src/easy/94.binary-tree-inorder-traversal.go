package easy

/*
 * @lc app=leetcode id=94 lang=golang
 *
 * [94] Binary Tree Inorder Traversal
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
func inorderTraversal(root *TreeNode) []int {
	result := []int{}
	result = inorderTraversalCalc(root, result)
	return result
}

func inorderTraversalCalc(root *TreeNode, data []int) []int {
	if root == nil {
		return data
	}
	data = inorderTraversalCalc(root.Left, data)
	data = append(data, root.Val)
	data = inorderTraversalCalc(root.Right, data)
	return data
}

// @lc code=end
