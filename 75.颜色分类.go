/*
 * @lc app=leetcode.cn id=75 lang=golang
 *
 * [75] 颜色分类
 */

// @lc code=start
func sortColors(nums []int) {
	left, right := 0, len(nums)-1

	for i := 0; i <= right; i++ {
		for i <= right && nums[i] == 2 {
			nums[i], nums[right] = nums[right], nums[i]
			right--
		}

		if nums[i] == 0 {
			nums[i], nums[left] = nums[left], nums[i]
			left++
		}
	}
}

// @lc code=end

