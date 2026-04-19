/*
 * @lc app=leetcode.cn id=42 lang=golang
 *
 * [42] 接雨水
 */

// @lc code=start
func trap(height []int) int {
	if len(height) < 3 {
		return 0
	}
	left, right := 0, len(height)-1
	leftMax, rightMax := 0, 0
	res := 0
	for left < right {
		// 更新左右两边的最高记录
		if height[left] > leftMax {
			leftMax = height[left]
		}
		if height[right] > rightMax {
			rightMax = height[right]
		}

		// 哪边短，就结算哪边
		if leftMax < rightMax {
			res += leftMax - height[left]
			left++
		} else {
			res += rightMax - height[right]
			right--
		}
	}
	return res
}

// @lc code=end

