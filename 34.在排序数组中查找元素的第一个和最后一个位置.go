/*
 * @lc app=leetcode.cn id=34 lang=golang
 *
 * [34] 在排序数组中查找元素的第一个和最后一个位置
 */

// @lc code=start
func searchRange(nums []int, target int) []int {
	// 顺序数组,二分查找
	left := findBound(nums, target, true)
	right := findBound(nums, target, false)
	return []int{left, right}
}

// 辅助函数
func findBound(nums []int, target int, isLeft bool) int {
	low, high := 0, len(nums)-1
	res := -1

	for low <= high {
		mid := low + (high-low)/2
		if nums[mid] < target {
			low = mid + 1
		} else if nums[mid] > target {
			high = mid - 1
		} else {
			res = mid
			if isLeft {
				high = mid - 1
			} else {
				low = mid + 1
			}
		}
	}
	return res
}

// @lc code=end

