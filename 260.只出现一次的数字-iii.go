/*
 * @lc app=leetcode.cn id=260 lang=golang
 *
 * [260] 只出现一次的数字 III
 */

// @lc code=start
func singleNumber(nums []int) []int {
	temp := 0
	for i := 0; i < len(nums); i++ {
		temp ^= nums[i]
	}

	result := []int{temp, temp}
	// 去掉末尾的1后异或temp就得到最后一个1的位置
	temp = (temp & (temp - 1)) ^ temp
	for i := 0; i < len(nums); i++ {
		if temp&nums[i] == 0 {
			result[0] ^= nums[i]
		} else {
			result[1] ^= nums[i]
		}
	}
	return result
}

// @lc code=end

