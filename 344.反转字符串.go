/*
 * @lc app=leetcode.cn id=344 lang=golang
 *
 * [344] 反转字符串
 */

// @lc code=start
func reverseString(s []byte) {
	// 定义双指针
	left, right := 0, len(s)-1
	// 双指针同时向中间移动，直到相遇
	for left < right {
		// 交换左右指针所指的元素
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}
}

// @lc code=end

