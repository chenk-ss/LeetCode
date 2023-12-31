package medium

/*
 * @lc app=leetcode id=199 lang=golang
 *
 * [199] Binary Tree Right Side View
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
func rightSideView(root *TreeNode) []int {
	res := []int{}
	rightSideViewFunc(root, 0, &res)
	return res
}

func rightSideViewFunc(root *TreeNode, level int, res *[]int) {
	if root == nil {
		return
	}
	if len(*res) == level {
		*res = append(*res, 0)
	}
	(*res)[level] = root.Val
	rightSideViewFunc(root.Left, level+1, res)
	rightSideViewFunc(root.Right, level+1, res)
}

// @lc code=end
