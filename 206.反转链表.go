/*
 * @lc app=leetcode.cn id=206 lang=golang
 *
 * [206] 反转链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	dummy := new(ListNode)
	temp := new(ListNode)

	for head != nil {
		temp = head
		head = head.Next

		temp.Next = dummy.Next
		dummy.Next = temp
	}
	return dummy.Next
}

// @lc code=end

