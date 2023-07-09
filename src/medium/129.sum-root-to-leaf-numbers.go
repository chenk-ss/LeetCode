package medium

/*
 * @lc app=leetcode id=129 lang=golang
 *
 * [129] Sum Root to Leaf Numbers
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
func sumNumbers(root *TreeNode) int {
	return sumNumbersFunc(root, 0)
}

func sumNumbersFunc(root *TreeNode, sum int) int {
	if root == nil {
		return sum
	}
	if root.Left == nil && root.Right == nil {
		return sum*10 + root.Val
	}
	tmp := 0
	if root.Left != nil {
		tmp += sumNumbersFunc(root.Left, sum*10+root.Val)
	}
	if root.Right != nil {
		tmp += sumNumbersFunc(root.Right, sum*10+root.Val)
	}
	return tmp
}

// @lc code=end
