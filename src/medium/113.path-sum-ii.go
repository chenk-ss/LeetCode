package medium

/*
 * @lc app=leetcode id=113 lang=golang
 *
 * [113] Path Sum II
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
func pathSum113(root *TreeNode, targetSum int) [][]int {
	res := [][]int{}
	if root == nil {
		return res
	}
	if root.Left == nil && root.Right == nil && root.Val == targetSum {
		res = append(res, []int{root.Val})
		return res
	}
	for _, v := range pathSum113(root.Left, targetSum-root.Val) {
		res = append(res, append([]int{root.Val}, v...))
	}
	for _, v := range pathSum113(root.Right, targetSum-root.Val) {
		res = append(res, append([]int{root.Val}, v...))
	}
	return res
}

// @lc code=end
