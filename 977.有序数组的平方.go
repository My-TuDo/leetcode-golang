/*
 * @lc app=leetcode.cn id=977 lang=golang
 *
 * [977] 有序数组的平方
 */

// @lc code=start
func sortedSquares(nums []int) []int {
	// 定义一个新的切片来存储平方后的结果
	result := make([]int, len(nums))
	left, right := 0, len(nums)-1
	index := len(nums) - 1

	// 使用双指针从两端向中间遍历
	for left <= right {
		leftSquare := nums[left] * nums[left]
		rightSquare := nums[right] * nums[right]
		if leftSquare > rightSquare {
			result[index] = leftSquare
			left++
		} else {
			result[index] = rightSquare
			right--
		}
		index--
	}
	return result
}

// @lc code=end

