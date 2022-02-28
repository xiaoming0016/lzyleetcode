/*
 * @lc app=leetcode.cn id=83 lang=golang
 *
 * [83] 删除排序链表中的重复元素
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
	temp := head
	for temp != nil {
		if temp.Next != nil && temp.Val == temp.Next.Val {
			if temp.Next.Next != nil {
				temp.Next = temp.Next.Next
			} else {
				temp.Next = nil
			}
		}
		temp = temp.Next
	}

	return head
}

// @lc code=end

