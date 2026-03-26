/*
 * @lc app=leetcode.cn id=217 lang=golang
 *
 * [217] 存在重复元素
 */

// @lc code=start
func containsDuplicate(nums []int) bool {
	// 定义一个map来存储已经出现过的元素
	seen := map[int]struct{}{}
	for _, num := range nums {
		// 如果元素已经存在于map中，说明存在重复元素，返回true
		if _, ok := seen[num]; ok {
			return true
		}
		// 否则将元素添加到map中
		seen[num] = struct{}{}
	}
	return false
}

// @lc code=end

