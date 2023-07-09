package medium

import "math"

/*
 * @lc app=leetcode id=222 lang=golang
 *
 * [222] Count Complete Tree Nodes
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
func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	level := countNodesFunc(root, 1)
	count := countLastLevelNodeFunc(root, 1, level, 0)
	return count + int(math.Pow(2, float64(level)-1)) - 1
}

func countNodesFunc(root *TreeNode, level int) int {
	if root == nil {
		return level
	}
	node := root
	for node.Left != nil {
		level++
		node = node.Left
	}
	return level
}

func countLastLevelNodeFunc(root *TreeNode, level, totalLevel, count int) int {
	if root == nil {
		return count
	}
	if level == totalLevel {
		return count + 1
	}
	count = countLastLevelNodeFunc(root.Left, level+1, totalLevel, count)
	count = countLastLevelNodeFunc(root.Right, level+1, totalLevel, count)
	return count
}

// @lc code=end
