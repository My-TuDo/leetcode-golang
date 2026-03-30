/*
 * @lc app=leetcode.cn id=704 lang=golang
 *
 * [704] 二分查找
 */

// @lc code=start
func search(nums []int, target int) int {
	// 定义左右指针
	left, right := 0, len(nums)-1

	// 二分查找
	for left <= right {
		mid := left + (right-left)/2 // 计算中间索引，避免溢出
		if nums[mid] == target {
			return mid // 找到目标，返回索引
		} else if nums[mid] < target {
			left = mid + 1 // 目标在右半部分
		} else {
			right = mid - 1 // 目标在左半部分
		}
	}
	return -1 // 没有找到目标，返回 -1

}

// @lc code=end

