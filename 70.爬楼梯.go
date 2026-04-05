/*
 * @lc app=leetcode.cn id=70 lang=golang
 *
 * [70] 爬楼梯
 */

// @lc code=start
func climbStairs(n int) int {
	// 斐波那契数列
	if n <= 2 {
		return n
	}
	p, q, r := 1, 2, 0
	for i := 3; i <= n; i++ {
		r = p + q
		p = q
		q = r
	}
	return q
}

// @lc code=end

