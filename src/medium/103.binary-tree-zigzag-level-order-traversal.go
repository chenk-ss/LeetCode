package medium

/*
 * @lc app=leetcode id=103 lang=golang
 *
 * [103] Binary Tree Zigzag Level Order Traversal
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
func zigzagLevelOrder(root *TreeNode) [][]int {
	res := [][]int{}
	if root == nil {
		return res
	}
	zigzagLevelOrderFunc(root, 0, &res)
	return res
}

func zigzagLevelOrderFunc(root *TreeNode, level int, res *[][]int) {
	if root == nil {
		return
	}
	if len(*res) == level {
		*res = append(*res, []int{})
	}
	if level%2 == 0 {
		(*res)[level] = append((*res)[level], root.Val)
	} else {
		(*res)[level] = append([]int{root.Val}, (*res)[level]...)
	}
	zigzagLevelOrderFunc(root.Left, level+1, res)
	zigzagLevelOrderFunc(root.Right, level+1, res)
}

// @lc code=end
