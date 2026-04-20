/*
 * @lc app=leetcode.cn id=560 lang=golang
 *
 * [560] 和为 K 的子数组
 */

// @lc code=start
func subarraySum(nums []int, k int) int {
	count := 0
	preSum := 0
	// m 记录前缀和出现的次数：map[前缀和]=出现次数
	m := make(map[int]int)
	m[0] = 1

	for _, num := range nums {
		preSum += num

		if times, ok := m[preSum-k]; ok {
			count += times
		}

		m[preSum]++
	}
	return count
}

// @lc code=end

