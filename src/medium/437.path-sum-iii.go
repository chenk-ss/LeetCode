package medium

/*
 * @lc app=leetcode id=437 lang=golang
 *
 * [437] Path Sum III
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
func pathSum(root *TreeNode, targetSum int) int {
	return pathSumFunc437(root, targetSum, []int{})
}

func pathSumFunc437(root *TreeNode, targetSum int, list []int) int {
	if root == nil {
		return 0
	}
	list = append(list, root.Val)
	sum := 0
	count := 0
	for i := len(list) - 1; i >= 0; i-- {
		sum += list[i]
		if sum == targetSum {
			count++
		}
	}
	count += pathSumFunc437(root.Left, targetSum, list)
	count += pathSumFunc437(root.Right, targetSum, list)
	return count
}

// @lc code=end
