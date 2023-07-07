package medium

/*
 * @lc app=leetcode id=173 lang=golang
 *
 * [173] Binary Search Tree Iterator
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
type BSTIterator struct {
	vals []int
}

func Constructor173(root *TreeNode) BSTIterator {
	list := []int{}
	BSTIteratorFunc(root, &list)
	return BSTIterator{vals: list}
}

func BSTIteratorFunc(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}
	BSTIteratorFunc(root.Left, res)
	*res = append(*res, root.Val)
	BSTIteratorFunc(root.Right, res)
}

func (this *BSTIterator) Next() int {
	res := this.vals[0]
	this.vals = this.vals[1:]
	return res
}

func (this *BSTIterator) HasNext() bool {
	return len(this.vals) != 0
}

/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */
// @lc code=end
