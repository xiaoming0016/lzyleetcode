package main

import "fmt"

// 62、不同路径
func uniquePaths(m int, n int) int {
	if m <= 0 || n <= 0 {
		return 0
	}
	if m == 1 {
		return 1
	}
	if n == 1 {
		return 1
	}

	dp := [][]int{}
	for i := 0; i < m; i++ {
		dp = append(dp, make([]int, n))
	}
	for i := 0; i < n; i++ {
		dp[0][i] = 1
	}
	for i := 0; i < m; i++ {
		dp[i][0] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}

	return dp[m-1][n-1]
}

// 63，不同路径2
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

// 64、最短路径和
func minPathSum(grid [][]int) int {
	m := len(grid)
	if m == 0 {
		return 0
	}
	n := len(grid[0])
	if n == 0 {
		return 0
	}
	dp := [][]int{}
	for i := 0; i < m; i++ {
		dp = append(dp, make([]int, n))
	}
	dp[0][0] = grid[0][0]
	for i := 1; i < m; i++ {
		dp[i][0] = dp[i-1][0] + grid[i][0]
	}
	for j := 1; j < n; j++ {
		dp[0][j] = dp[0][j-1] + grid[0][j]
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
		}
	}
	fmt.Println(dp)
	return dp[m-1][n-1]
}

// 70、爬楼梯
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

func min(data1, data2 int) int {
	if data1 < data2 {
		return data1
	}
	return data2
}
