/*
 * @lc app=leetcode.cn id=121 lang=golang
 *
 * [121] 买卖股票的最佳时机
 */

// @lc code=start
func maxProfit(prices []int) int {
	slow := 0
	max := 0
	for fast := 0; fast < len(prices); fast++ {
		diff := prices[fast] - prices[slow]
		if diff < 0 {
			slow = fast
		} else if diff > max {
			max = diff
		}
	}
	return max
}

// @lc code=end

