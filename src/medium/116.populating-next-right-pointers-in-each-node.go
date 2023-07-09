package medium

/*
 * @lc app=leetcode id=116 lang=golang
 *
 * [116] Populating Next Right Pointers in Each Node
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */

func connect116(root *Node) *Node {
	res := [][]*Node{}
	connectFunc116(root, &res, 0)
	for _, v := range res {
		for i := 0; i < len(v)-1; i++ {
			v[i].Next = v[i+1]
		}
	}
	return root
}

func connectFunc116(root *Node, res *[][]*Node, level int) {
	if root == nil {
		return
	}
	if len(*res) == level {
		*res = append(*res, []*Node{})
	}
	(*res)[level] = append((*res)[level], root)
	connectFunc116(root.Left, res, level+1)
	connectFunc116(root.Right, res, level+1)
}

// @lc code=end
