package medium

/*
 * @lc app=leetcode id=133 lang=golang
 *
 * [133] Clone Graph
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */

func cloneGraph(node *GraphNode) *GraphNode {
	if node == nil {
		return nil
	}
	m := make(map[*GraphNode]*GraphNode)
	var cloneGraphFunc func(onode *GraphNode) *GraphNode
	cloneGraphFunc = func(onode *GraphNode) *GraphNode {
		if v, ok := m[onode]; ok {
			return v
		}
		clone := &GraphNode{Val: onode.Val, Neighbors: []*GraphNode{}}
		m[onode] = clone
		for _, v := range onode.Neighbors {
			clone.Neighbors = append(clone.Neighbors, cloneGraphFunc(v))
		}
		return clone
	}
	return cloneGraphFunc(node)
}

// @lc code=end
