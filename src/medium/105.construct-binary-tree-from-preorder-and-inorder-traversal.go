package medium

/*
 * @lc app=leetcode id=105 lang=golang
 *
 * [105] Construct Binary Tree from Preorder and Inorder Traversal
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
func buildTree105(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}
	root := &TreeNode{}
	root.Val = preorder[0]
	if len(preorder) == 1 || len(inorder) == 1 {
		return root
	}
	index := 0
	for i := range inorder {
		if inorder[i] == preorder[0] {
			index = i
			break
		}
	}
	root.Left = buildTree(preorder[1:1+index], inorder[:index])
	root.Right = buildTree(preorder[1+index:], inorder[index+1:])
	return root
}

// @lc code=end
