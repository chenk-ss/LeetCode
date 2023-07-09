package medium

/*
 * @lc app=leetcode id=138 lang=golang
 *
 * [138] Copy List with Random Pointer
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList(head *RandomNode) *RandomNode {
	m := make(map[*RandomNode]*RandomNode)
	var copyRandomListFunc func(node *RandomNode) *RandomNode
	copyRandomListFunc = func(node *RandomNode) *RandomNode {
		if node == nil {
			return nil
		}
		if v, ok := m[node]; ok {
			return v
		}
		clone := &RandomNode{Val: node.Val}
		m[node] = clone
		clone.Next = copyRandomListFunc(node.Next)
		clone.Random = copyRandomListFunc(node.Random)
		return clone
	}
	return copyRandomListFunc(head)
}

// @lc code=end
