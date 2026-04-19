/*
 * @lc app=leetcode.cn id=62 lang=golang
 *
 * [62] 不同路径
 */

// @lc code=start
func uniquePaths(m int, n int) int {
	// 1. 创建二维 DP 数组
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	// 2. 初始化边界
	for i := 0; i < m; i++ {
		dp[i][0] = 1
	} // 第一列
	for j := 0; j < n; j++ {
		dp[0][j] = 1
	} // 第一行

	// 动态规划填表
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}

	return dp[m-1][n-1]
}

// @lc code=end

