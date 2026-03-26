/*
 * @lc app=leetcode.cn id=136 lang=golang
 *
 * [136] 只出现一次的数字
 */

// @lc code=start
func singleNumber(nums []int) int {
	// 使用异或运算，任何数与自己异或结果为0，任何数与0异或结果为自己
	result := 0
	for _, num := range nums {
		result ^= num
	}
	return result
}

// @lc code=end

