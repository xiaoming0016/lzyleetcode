/*
 * @lc app=leetcode.cn id=63 lang=golang
 *
 * [63] 不同路径 II
 */

// @lc code=start

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m := len(obstacleGrid)
	if m == 0 {
		return 0
	}
	n := len(obstacleGrid[0])
	if n == 0 {
		return 0
	}
	if m == 1 {
		for i := 0; i < n; i++ {
			if obstacleGrid[0][i] == 1 {
				return 0
			}
		}
		return 1
	}
	if n == 1 {
		for i := 0; i < m; i++ {
			if obstacleGrid[i][0] == 1 {
				return 0
			}
		}
		return 1
	}
	if obstacleGrid[0][0] == 1 {
		return 0
	}
	dp := [][]int{}
	for i := 0; i < m; i++ {
		dp = append(dp, make([]int, n))
	}
	dp[0][0] = 1
	for i := 1; i < n; i++ {
		if obstacleGrid[0][i] == 0 && dp[0][i-1] == 1 {
			dp[0][i] = 1
		} else {
			dp[0][i] = 0
		}
	}
	for i := 1; i < m; i++ {
		if obstacleGrid[i][0] == 0 && dp[i-1][0] == 1 {
			dp[i][0] = 1
		} else {
			dp[i][0] = 0
		}
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if obstacleGrid[i][j] == 0 {
				if dp[i-1][j] != 0 {
					dp[i][j] += dp[i-1][j]
				}
				if dp[i][j-1] != 0 {
					dp[i][j] += dp[i][j-1]
				}
			} else {
				dp[i][j] = 0
			}
		}
	}
	fmt.Println(dp)
	return dp[m-1][n-1]
}

// @lc code=end

