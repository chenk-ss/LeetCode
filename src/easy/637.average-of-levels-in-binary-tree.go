package easy

/*
 * @lc app=leetcode id=637 lang=golang
 *
 * [637] Average of Levels in Binary Tree
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
func averageOfLevels(root *TreeNode) []float64 {
	dp := [][]int{}
	averageOfLevelsFunc(root, 0, &dp)
	res := []float64{}
	for i := 0; i < len(dp); i++ {
		res = append(res, float64(dp[i][0])/float64(dp[i][1]))
	}
	return res
}

func averageOfLevelsFunc(root *TreeNode, level int, dp *[][]int) {
	if root == nil {
		return
	}
	if len(*dp) == level {
		*dp = append(*dp, []int{0, 0})
	}
	(*dp)[level][0] += root.Val
	(*dp)[level][1]++
	averageOfLevelsFunc(root.Left, level+1, dp)
	averageOfLevelsFunc(root.Right, level+1, dp)
}

// @lc code=end
