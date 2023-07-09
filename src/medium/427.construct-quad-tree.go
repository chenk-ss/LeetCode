package medium

/*
 * @lc app=leetcode id=427 lang=golang
 *
 * [427] Construct Quad Tree
 */

// @lc code=start
/**
 * Definition for a QuadTree QuadNode.
 * type Node struct {
 *     Val bool
 *     IsLeaf bool
 *     TopLeft *Node
 *     TopRight *Node
 *     BottomLeft *Node
 *     BottomRight *Node
 * }
 */

func construct(grid [][]int) *QuadNode {
	var constructFunc func(top, left, width int) *QuadNode
	constructFunc = func(top, left, width int) *QuadNode {
		flag := true
		val := grid[top][left]
		for i := top; i < top+width; i++ {
			for j := left; j < left+width; j++ {
				if grid[i][j] != val {
					flag = false
				}
			}
		}
		if flag {
			return &QuadNode{val == 1, true, nil, nil, nil, nil}
		}
		w := width / 2
		return &QuadNode{
			Val:         val == 1,
			IsLeaf:      flag,
			TopLeft:     constructFunc(top, left, w),
			TopRight:    constructFunc(top, left+w, w),
			BottomLeft:  constructFunc(top+w, left, w),
			BottomRight: constructFunc(top+w, left+w, w),
		}
	}
	return constructFunc(0, 0, len(grid))
}

// @lc code=end
