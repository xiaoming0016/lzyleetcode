package main

import "fmt"

// 链表
type ListNode struct {
	Val  int
	Next *ListNode
}

// 打印链表
func printList(head *ListNode) {
	temp := head

	for i := 0; temp != nil; i++ {
		fmt.Print(temp.Val, " ")
		temp = temp.Next
	}
	fmt.Print("\n")
}

// 生成链表
func createList(list []int) *ListNode {
	if len(list) == 0 {
		return &ListNode{Val: 0, Next: nil}
	}

	head := new(ListNode)
	temp := head
	for i := 0; i < len(list); i++ {
		node := &ListNode{Val: list[i], Next: nil}
		temp.Next = node
		temp = node
	}

	return head.Next
}

// func main() {
// list := []int{1, 4, 3, 2, 5, 2}
// head := createList(list)
// printList(head)

// head = partition(head, 3)
// printList(head)

// 	var items []int = []int{10, 20, 30, 40}
// 	fmt.Println(items)
// 	for _, item := range items {
// 		item *= 2
// 	}
// 	fmt.Println(items)
// }

// 删除所有重复元素
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

// 翻转链表
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

// 翻转链表的从left到right
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

// 合并两个有序列表
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

// 旋转链表
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

// 分割链表
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

// 奇偶链表
// func oddEvenList(head *ListNode) *ListNode {
// 	if head == nil {
// 		return nil
// 	}
// 	temp := head
// 	for i := 1; temp != nil; i++ {
// 		if i%2 == 0 {

// 		}
// 	}
// }

// 两数相加
func listToNum(list *ListNode) int {

	return 0
}
func numToList(list *ListNode) *ListNode {
	return nil
}
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := new(ListNode)

	num1 := listToNum(l1)
	num2 := listToNum(l2)
	fmt.Println(num1, num2)

	return head
}
