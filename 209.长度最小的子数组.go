/*
 * @lc app=leetcode.cn id=209 lang=golang
 *
 * [209] 长度最小的子数组
 */

// @lc code=start
func minSubArrayLen(target int, nums []int) int {
	// 滑动窗口
	left, right := 0, 0     // 左右指针维护窗口边界
	minLen := math.MaxInt32 // 最小的长度
	sum := 0                // 记录和

	for right < len(nums) {
		sum += nums[right]

		for sum >= target {
			// 每次收缩时都更新一下最小值
			currentLen := right - left + 1
			if minLen > currentLen {
				minLen = currentLen
			}
			sum -= nums[left]
			left++
		}
		right++
	}
	if minLen == math.MaxInt32 {
		return 0
	}
	return minLen
}

// @lc code=end

