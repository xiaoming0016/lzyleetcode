/*
 * @lc app=leetcode.cn id=82 lang=golang
 *
 * [82] 删除排序链表中的重复元素 II
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	dummy := new(ListNode)
	dummy.Next = head

	head = dummy
	for head.Next != nil && head.Next.Next != nil {
		if head.Next.Val == head.Next.Next.Val {
			v := head.Next.Val
			for head.Next != nil && v == head.Next.Val {
				head.Next = head.Next.Next
			}
		} else {
			head = head.Next

		}
	}
	return dummy.Next
}

// @lc code=end

