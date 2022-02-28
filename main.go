package main

import "fmt"

func testTree() {
	tree := createTree()
	preorderTraversal(tree)
	res := inorderTraversal(tree)
	fmt.Println(res)
	nums := []int{1, 2, 1, 3, 2, 4}
	temp := 0
	for i := 0; i < len(nums); i++ {
		temp ^= nums[i]
	}

	result := []int{temp, temp}
	fmt.Println(temp)
	// 去掉末尾的1后异或temp就得到最后一个1的位置
	fmt.Println(temp&(temp-1), temp&(temp-1)^temp)

	temp = (temp & (temp - 1)) ^ temp
	for i := 0; i < len(nums); i++ {
		if temp&nums[i] == 0 {
			result[0] ^= nums[i]
		} else {
			result[1] ^= nums[i]
		}
	}
	fmt.Println(6 & 5)
	fmt.Println((6 & 5) ^ 6)
}

func testList() {
	list := []int{1, 4, 3, 2, 5, 2}
	head := createList(list)
	printList(head)

	head = partition(head, 3)
	printList(head)

	var items []int = []int{10, 20, 30, 40}
	fmt.Println(items)
	for _, item := range items {
		item *= 2
	}
	fmt.Println(items)
}

func testDynaMicProgram() {

}
func main() {
	// a := [][]int{{0, 1, 0, 0, 0}, {1, 0, 0, 0, 0}, {0, 0, 0, 0, 0}, {0, 0, 0, 0, 0}}
	// res := uniquePathsWithObstacles(a)
	// fmt.Println(res)
	a := [][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}
	res := minPathSum(a)
	fmt.Println(res)
}
