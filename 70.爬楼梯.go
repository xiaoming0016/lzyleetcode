/*
 * @lc app=leetcode.cn id=70 lang=golang
 *
 * [70] 爬楼梯
 */

// @lc code=start
func climbStairs(n int) int {
	if n == 0 || n == 1 || n == 2 {
		return n
	}
	dp := []int{}
	dp = append(dp, 0)
	dp = append(dp, 1)
	dp = append(dp, 2)
	for i := 3; i <= n; i++ {
		dp = append(dp, dp[i-1]+dp[i-2])
	}
	return dp[n]
}

// @lc code=end

