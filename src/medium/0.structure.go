package medium

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type ListNode struct {
	Val  int
	Next *ListNode
}

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

type RandomNode struct {
	Val    int
	Next   *RandomNode
	Random *RandomNode
}

type GraphNode struct {
	Val       int
	Neighbors []*GraphNode
}

type QuadNode struct {
	Val         bool
	IsLeaf      bool
	TopLeft     *QuadNode
	TopRight    *QuadNode
	BottomLeft  *QuadNode
	BottomRight *QuadNode
}

type NTreeNode struct {
	Val      int
	Children []*NTreeNode
}
