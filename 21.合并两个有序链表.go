/*
 * @lc app=leetcode.cn id=21 lang=golang
 *
 * [21] 合并两个有序链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}

	newHead := new(ListNode)
	ptr1 := list1
	ptr2 := list2
	temp := newHead
	for ptr1 != nil && ptr2 != nil {
		if ptr1.Val <= ptr2.Val {
			temp.Next = ptr1
			ptr1 = ptr1.Next
		} else {
			temp.Next = ptr2
			ptr2 = ptr2.Next
		}
		temp = temp.Next
	}

	for ptr1 != nil {
		temp.Next = ptr1
		ptr1 = ptr1.Next
		temp = temp.Next
	}

	for ptr2 != nil {
		temp.Next = ptr2
		ptr2 = ptr2.Next
		temp = temp.Next
	}

	return newHead.Next
}

// @lc code=end

