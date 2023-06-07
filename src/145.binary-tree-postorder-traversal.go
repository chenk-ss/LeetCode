package src

/*
 * @lc app=leetcode id=145 lang=golang
 *
 * [145] Binary Tree Postorder Traversal
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
func postorderTraversal(root *TreeNode) []int {
	return postorderTraversalFunc(root, []int{})
}

func postorderTraversalFunc(root *TreeNode, list []int) []int {
	if root == nil {
		return list
	}
	list = postorderTraversalFunc(root.Left, list)
	list = postorderTraversalFunc(root.Right, list)
	list = append(list, root.Val)
	return list
}

// @lc code=end
