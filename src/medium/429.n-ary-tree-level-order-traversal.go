package medium

/*
 * @lc app=leetcode id=429 lang=golang
 *
 * [429] N-ary Tree Level Order Traversal
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func levelOrder(root *NTreeNode) [][]int {
	res := [][]int{}
	levelOrderFunc429(root, &res, 0)
	return res
}

func levelOrderFunc429(root *NTreeNode, res *[][]int, level int) {
	if root == nil {
		return
	}
	if len(*res) == level {
		*res = append(*res, []int{})
	}
	(*res)[level] = append((*res)[level], root.Val)
	for _, v := range root.Children {
		levelOrderFunc429(v, res, level+1)
	}
}

// @lc code=end
