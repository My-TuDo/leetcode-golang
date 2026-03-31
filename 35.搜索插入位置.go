/*
 * @lc app=leetcode.cn id=35 lang=golang
 *
 * [35] 搜索插入位置
 */

// @lc code=start
func searchInsert(nums []int, target int) int {
	// 二分查找
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid // 找到目标，返回索引
		} else if nums[mid] < target {
			left = mid + 1 // 目标在右半部分
		} else {
			right = mid - 1 // 目标在左半部分
		}
	}
	return left // 没有找到目标，返回插入位置的索引
}

// @lc code=end

