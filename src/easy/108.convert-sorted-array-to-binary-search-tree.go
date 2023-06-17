package easy

/*
 * @lc app=leetcode id=108 lang=golang
 *
 * [108] Convert Sorted Array to Binary Search Tree
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
func sortedArrayToBST(nums []int) *TreeNode {
	length := len(nums)
	if length == 0 {
		return nil
	}
	root := new(TreeNode)
	mid := length / 2
	root.Val = nums[mid]
	root.Left = sortedArrayToBST(nums[0:mid])
	root.Right = sortedArrayToBST(nums[mid+1 : length])
	return root
}

// @lc code=end
