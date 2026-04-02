/*
 * @lc app=leetcode.cn id=3 lang=golang
 *
 * [3] 无重复字符的最长子串
 */

// @lc code=start
func lengthOfLongestSubstring(s string) int {
	// 使用数组
	data := []byte(s) // 转换为byte数组
	n := len(data)
	if n <= 1 {
		return n
	}

	lastPos := [128]int{} // 用于记录元素出现的最新索引，
	for i := range lastPos {
		lastPos[i] = -1
	}

	maxLen := 0 // 记录最大长度
	start := 0  //窗口左边界
	for i := 0; i < n; i++ {
		char := data[i]                  // 遍历获取字符串s
		if lastPos[int(char)] >= start { // 如果下标大于等于start，说明字符重复出现
			start = lastPos[char] + 1 // 将上一个重复字符之前的字符串全部剔除
		}
		lastPos[int(char)] = i  // 更新字符索引
		if i-start+1 > maxLen { // 更新最大长度
			maxLen = i - start + 1
		}
	}
	return maxLen
}

// @lc code=end

