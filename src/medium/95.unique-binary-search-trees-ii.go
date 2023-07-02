package medium

/*
 * @lc app=leetcode id=95 lang=golang
 *
 * [95] Unique Binary Search Trees II
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
func generateTrees(n int) []*TreeNode {
	return generateTreesFunc(1, n)
}

func generateTreesFunc(start, end int) []*TreeNode {
	if start > end {
		return []*TreeNode{nil}
	}
	res := []*TreeNode{}
	for i := start; i <= end; i++ {
		leftTree := generateTreesFunc(start, i-1)
		rightTree := generateTreesFunc(i+1, end)
		for _, left := range leftTree {
			for _, right := range rightTree {
				res = append(res, &TreeNode{Val: i, Left: left, Right: right})
			}
		}
	}
	return res
}

// @lc code=end
