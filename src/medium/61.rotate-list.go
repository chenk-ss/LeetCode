package medium

/*
 * @lc app=leetcode id=61 lang=golang
 *
 * [61] Rotate List
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || k == 0 {
		return head
	}
	nHead := ListNode{-1, head}
	slow := &nHead
	fast := &nHead
	for i := 0; i < k; i++ {
		fast = fast.Next
		if fast == nil {
			fast = nHead.Next
		}
	}
	for {
		if fast.Next == nil {
			prev := slow.Next
			slow.Next = nil
			fast.Next = nHead.Next
			nHead.Next = prev
			break
		}
		slow = slow.Next
		fast = fast.Next
	}
	return nHead.Next
}

// @lc code=end
