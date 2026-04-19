/*
 * @lc app=leetcode.cn id=322 lang=golang
 *
 * [322] 零钱兑换
 */

// @lc code=start
func coinChange(coins []int, amount int) int {
	// 1. 初始化 DP 数组，初始值为 amuout + 1
	dp := make([]int, amount+1)
	for i := range dp {
		dp[i] = amount + 1 // 最大值占位
	}
	dp[0] = 0

	// 2. 遍历每一个金额
	for i := 1; i <= amount; i++ {
		// 3. 尝试每一枚硬币
		for _, coin := range coins {
			if i-coin >= 0 {
				// 状态转移： 当前最少硬币数 vs 选了这枚硬币后的新总数
				dp[i] = min(dp[i], dp[i-coin]+1)
			}
		}
	}

	// 判断无解情况
	if dp[amount] > amount {
		return -1
	}
	return dp[amount]
}

// @lc code=end

