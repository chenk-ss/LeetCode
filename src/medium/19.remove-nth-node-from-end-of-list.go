package medium

/*
 * @lc app=leetcode id=19 lang=golang
 *
 * [19] Remove Nth Node From End of List
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	newHead := &ListNode{-1, head}
	node1, node2 := newHead, newHead
	for node1.Next != nil {
		if n <= 0 {
			node2 = node2.Next
		}
		node1 = node1.Next
		n--
	}
	node2.Next = node2.Next.Next
	return newHead.Next
}

func removeNthFromEnd2(head *ListNode, n int) *ListNode {
	length, node := 0, head
	for node != nil {
		length++
		node = node.Next
	}
	newHead := head
	count := 0
	if length == n {
		return head.Next
	}
	for newHead != nil {
		if count == length-n-1 {
			if newHead.Next != nil {
				newHead.Next = newHead.Next.Next
			} else {
				newHead.Next = nil
			}
			return head
		}
		count++
		newHead = newHead.Next
	}

	return head
}

// @lc code=end
