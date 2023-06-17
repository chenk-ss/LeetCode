package easy

/*
 * @lc app=leetcode id=234 lang=golang
 *
 * [234] Palindrome Linked List
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func LinkedListIsPalindrome(head *ListNode) bool {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	var lastHalfNodeReverse *ListNode
	for slow != nil {
		next := slow.Next
		slow.Next = lastHalfNodeReverse
		lastHalfNodeReverse = slow
		slow = next
	}
	for lastHalfNodeReverse != nil && head != nil {
		if lastHalfNodeReverse.Val != head.Val {
			return false
		}
		lastHalfNodeReverse = lastHalfNodeReverse.Next
		head = head.Next
	}
	return true
}

func LinkedListIsPalindrome2(head *ListNode) bool {
	list := []int{}
	for head != nil {
		list = append(list, head.Val)
		head = head.Next
	}
	for i := 0; i < len(list)/2; i++ {
		if list[i] != list[len(list)-1-i] {
			return false
		}
	}
	return true
}

// @lc code=end
