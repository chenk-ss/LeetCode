package medium

/*
 * @lc app=leetcode id=117 lang=golang
 *
 * [117] Populating Next Right Pointers in Each Node II
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

func connect117(root *Node) *Node {
	res := [][]*Node{}
	connectFunc117(root, &res, 0)
	for _, v := range res {
		for i := 0; i < len(v)-1; i++ {
			v[i].Next = v[i+1]
		}
	}
	return root
}

func connectFunc117(root *Node, res *[][]*Node, level int) {
	if root == nil {
		return
	}
	if len(*res) == level {
		*res = append(*res, []*Node{})
	}
	(*res)[level] = append((*res)[level], root)
	connectFunc117(root.Left, res, level+1)
	connectFunc117(root.Right, res, level+1)
}

// @lc code=end
