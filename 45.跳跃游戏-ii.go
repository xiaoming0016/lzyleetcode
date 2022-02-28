/*
 * @lc app=leetcode.cn id=45 lang=golang
 *
 * [45] 跳跃游戏 II
 */

// @lc code=start

func jump(nums []int) int {
	if len(nums) == 0 || len(nums) == 1 {
		return 0
	}
	sum := 0
	count := 0
	length := len(nums)
	value := 0
	index := 0
	for i := 0; i < length-1; {
		sum++

		count = nums[i]
		value = 0
		for j := 1; j <= count; j++ {
			if (i + j) >= length {
				return sum
			}

			if value == 0 {
				value = nums[i+j]
				index = i + j
			} else {
				if nums[i+j] > value {
					value = nums[i+j]
					index = i + j
				}
			}
		}
		i = index
	}

	return sum
}

// @lc code=end

