package medium

/*
 * @lc app=leetcode id=230 lang=golang
 *
 * [230] Kth Smallest Element in a BST
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
func kthSmallest(root *TreeNode, k int) int {
	count, res := 0, 0
	var kthSmallestFunc func(root *TreeNode)
	kthSmallestFunc = func(root *TreeNode) {
		if root == nil {
			return
		}
		kthSmallestFunc(root.Left)
		count++
		if k == count {
			res = root.Val
			return
		}
		kthSmallestFunc(root.Right)
	}
	kthSmallestFunc(root)
	return res
}

// @lc code=end
