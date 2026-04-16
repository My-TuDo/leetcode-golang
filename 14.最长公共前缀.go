/*
 * @lc app=leetcode.cn id=14 lang=golang
 *
 * [14] 最长公共前缀
 */

// @lc code=start
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	first := strs[0]
	for i := 0; i < len(first); i++ {
		char := first[i]
		for j := 1; j < len(strs); j++ {
			if i == len(strs[j]) || strs[j][i] != char {
				return first[:i]
			}
		}
	}
	return first
}

// @lc code=end

