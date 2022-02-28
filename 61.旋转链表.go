/*
 * @lc app=leetcode.cn id=61 lang=golang
 *
 * [61] 旋转链表
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
	if k == 0 {
		return head
	}
	if head == nil {
		return nil
	}
	temp := new(ListNode)
	count := 0
	tail := new(ListNode)
	for temp = head; temp != nil; {
		count++
		if temp.Next != nil {
			tail = temp.Next
		}
		temp = temp.Next
	}
	if k == count {
		return head
	}

	if k > count {
		k = k % count
	}
	newHead := new(ListNode)
	temp = head
	for i := 0; i < k; i++ {
		temp = temp.Next
	}
	newHead.Next = temp
	tail.Next = head
	return newHead
}

// @lc code=end

