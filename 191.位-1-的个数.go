/*
 * @lc app=leetcode.cn id=191 lang=golang
 *
 * [191] 位1的个数
 */

// @lc code=start
func hammingWeight(n int) int {
	// 位运算
	count := 0
	for n > 0 {
		n &= (n - 1) // 每次消除最右边的1
		count++
	}
	return count
}

// @lc code=end

