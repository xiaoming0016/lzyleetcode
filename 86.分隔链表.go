/*
 * @lc app=leetcode.cn id=86 lang=golang
 *
 * [86] 分隔链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func partition(head *ListNode, x int) *ListNode {
	if head == nil {
		return nil
	}
	firstHead := new(ListNode)
	secondHead := new(ListNode)
	temp1, temp2 := firstHead, secondHead

	temp := head
	for temp != nil {
		if temp.Val >= x {
			temp2.Next = temp
			temp2 = temp2.Next
		} else {
			temp1.Next = temp
			temp1 = temp1.Next
		}
		temp = temp.Next
	}
	temp2.Next = nil
	temp1.Next = secondHead.Next
	return firstHead.Next
}

// @lc code=end

