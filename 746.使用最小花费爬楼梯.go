/*
 * @lc app=leetcode.cn id=746 lang=golang
 *
 * [746] 使用最小花费爬楼梯
 */

// @lc code=start
func minCostClimbingStairs(cost []int) int {
	n := len(cost)
	prev, curr := 0, 0
	for i := 2; i <= n; i++ {
		next := min(curr+cost[i-1], prev+cost[i-2])
		prev = curr
		curr = next
	}
	return curr
}

// @lc code=end

