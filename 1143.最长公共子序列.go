/*
 * @lc app=leetcode.cn id=1143 lang=golang
 *
 * [1143] 最长公共子序列
 */

// @lc code=start
func longestCommonSubsequence(text1 string, text2 string) int {
	m, n := len(text1), len(text2)
	// 1. 创建 dp 数组， 多留出第0行和第0列作为边界（初始值默认为0）
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	// 2. 填表
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			// 因为 dp 索引比字符串索引多1，所以对比字符是要用i-1和j-1
			if text1[i-1] == text2[j-1] {
				// 字符匹配： 左上方对角线的值加一
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				// 字符不匹配： 取左边或者上边的最大值
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[m][n]
}

// @lc code=end

