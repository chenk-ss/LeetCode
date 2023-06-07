package src

/*
 * @lc app=leetcode id=144 lang=golang
 *
 * [144] Binary Tree Preorder Traversal
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
func preorderTraversal(root *TreeNode) []int {
	return preorderTraversalFunc(root, []int{})
}

func preorderTraversalFunc(root *TreeNode, list []int) []int {
	if root == nil {
		return list
	}
	list = append(list, root.Val)
	list = preorderTraversalFunc(root.Left, list)
	list = preorderTraversalFunc(root.Right, list)
	return list
}

// @lc code=end
