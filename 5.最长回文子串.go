/*
 * @lc app=leetcode.cn id=5 lang=golang
 *
 * [5] 最长回文子串
 */

// @lc code=start
func longestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}

	start, end := 0, 0

	for i := 0; i < len(s); i++ {
		// 1. 假设回文是奇数长度，中心是 s[i]
		len1 := expandAroundCenter(s, i, i)
		// 2. 假设回文是偶数长度，中心是 s[i] 和 s[i+1]
		len2 := expandAroundCenter(s, i, i+1)

		// 区两次扩展中最长的那个
		maxLen := max(len1, len2)

		// 如果发现比之前记录的更长，更新起始和结束位置
		if maxLen > end-start {
			// 根据中心 i 和长度 maxLen 算出左右边界
			start = i - (maxLen-1)/2
			end = i + maxLen/2
		}
	}
	return s[start : end+1]
}

// 辅助函数： 向两边扩散， 返回扩散后的长度
func expandAroundCenter(s string, left, right int) int {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}

	// 退出循环时， s[left] != s[right], 所以长度是(right-1) - (left+1)+1
	return right - left - 1
}

// @lc code=end

