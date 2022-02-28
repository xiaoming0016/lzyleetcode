package main

import "fmt"

var nums = []int{1, 2, 0, 0, 3}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 先序遍历，递归
func preorderTraversal(root *TreeNode) {
	// if root == nil {
	// 	return
	// }

	// fmt.Println(root.Val)
	// preorderTraversal(root.Left)
	// preorderTraversal(root.Right)
	if root == nil {
		return
	}
	fmt.Println(root.Val)
	preorderTraversal(root.Left)
	preorderTraversal(root.Right)
}

// 先序遍历，非递归
func preorderTraversal2(root *TreeNode) []int {

	if root == nil {
		return nil
	}
	result := make([]int, 0)
	stack := make([]*TreeNode, 0)

	for root != nil || len(stack) != 0 {
		for root != nil {
			// 前序遍历，所以先保存结果
			result = append(result, root.Val)
			stack = append(stack, root)
			root = root.Left
		}
		// pop
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		root = node.Right
	}
	return result
}

// 中序遍历，非递归
func inorderTraversal(root *TreeNode) []int {
	result := make([]int, 0)
	if root == nil {
		return result
	}
	stack := make([]*TreeNode, 0)
	for len(stack) > 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left // 一直向左
		}
		// 弹出
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result = append(result, node.Val)
		root = node.Right
	}
	return result
}

// 后续遍历，非递归
func beforderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	result := make([]int, 0)
	stack := make([]*TreeNode, 0)
	var lastVisit *TreeNode
	for root != nil || len(stack) != 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		// 这里先看看，先不弹出
		node := stack[len(stack)-1]
		// 根节点必须在右节点弹出之后，再弹出
		if node.Right == nil || node.Right == lastVisit {
			stack = stack[:len(stack)-1] // pop
			result = append(result, node.Val)
			// 标记当前这个节点已经弹出过
			lastVisit = node
		} else {
			root = node.Right
		}
	}
	return result
}
func createTree() *TreeNode {
	if len(nums) > 0 {
		value := nums[0]
		nums = nums[1:]

		if value != 0 {
			node := &TreeNode{Val: value}
			node.Left = createTree()
			node.Right = createTree()
			return node
		}
	}
	return nil
}

// func main() {
// tree := createTree()
// preorderTraversal(tree)
// res := inorderTraversal(tree)
// fmt.Println(res)
// nums := []int{1, 2, 1, 3, 2, 4}
// temp := 0
// for i := 0; i < len(nums); i++ {
// 	temp ^= nums[i]
// }

// result := []int{temp, temp}
// fmt.Println(temp)
// // 去掉末尾的1后异或temp就得到最后一个1的位置
// fmt.Println(temp&(temp-1), temp&(temp-1)^temp)

// temp = (temp & (temp - 1)) ^ temp
// for i := 0; i < len(nums); i++ {
// 	if temp&nums[i] == 0 {
// 		result[0] ^= nums[i]
// 	} else {
// 		result[1] ^= nums[i]
// 	}
// }
// fmt.Println(6 & 5)
// fmt.Println((6 & 5) ^ 6)
// }

// func inorderTraversal(root *TreeNode) []int {
// 	return []int{0}
// }
