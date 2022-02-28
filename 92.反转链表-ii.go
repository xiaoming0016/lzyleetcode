/*
 * @lc app=leetcode.cn id=92 lang=golang
 *
 * [92] 反转链表 II
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func doReverse(head *ListNode, count int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	dummy := new(ListNode)
	temp := new(ListNode)
	dummytail := head

	for i := 0; i < count && head != nil; i++ {
		temp = head
		head = head.Next

		temp.Next = dummy.Next
		dummy.Next = temp
	}
	// tail.Next = head
	dummytail.Next = head
	return dummy.Next
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	if left == right {
		return head
	}
	dummy := new(ListNode)
	dummy.Next = head

	firstTail := dummy
	for i := 1; i < left && head != nil; i++ {
		firstTail = head
		head = head.Next
	}
	nextHead := doReverse(head, right-left+1)
	if dummy == firstTail {
		return nextHead
	}
	firstTail.Next = nextHead
	return dummy.Next
}

// @lc code=end

