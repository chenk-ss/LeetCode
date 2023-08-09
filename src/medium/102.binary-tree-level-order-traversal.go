package medium

/*
 * @lc app=leetcode id=102 lang=golang
 *
 * [102] Binary Tree Level Order Traversal
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
func levelOrder102(root *TreeNode) [][]int {
	res := [][]int{}
	levelOrderFunc(root, &res, 0)
	return res
}

func levelOrderFunc(root *TreeNode, res *[][]int, level int) {
	if root == nil {
		return
	}
	if len(*res) == level {
		*res = append(*res, []int{})
	}
	(*res)[level] = append((*res)[level], root.Val)
	levelOrderFunc(root.Left, res, level+1)
	levelOrderFunc(root.Right, res, level+1)
}

// @lc code=end
