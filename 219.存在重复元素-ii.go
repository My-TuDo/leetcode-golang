/*
 * @lc app=leetcode.cn id=219 lang=golang
 *
 * [219] 存在重复元素 II
 */

// @lc code=start
func containsNearbyDuplicate(nums []int, k int) bool {
	// 哈希表 元素：索引，遇到一个存入索引，相同的比较是否小于等于k，如果是，返回，不是，更新索引
	m := make(map[int]int, k)

	for i, n := range nums {
		if index, ok := m[n]; ok && i-index <= k {
			return true
		}
		// 存入操作
		m[n] = i
		// 滑动窗口
		if len(m) > k {
			delete(m, nums[i-k])
		}
	}
	return false
}

// @lc code=end

