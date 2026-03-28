/*
 * @lc app=leetcode.cn id=387 lang=golang
 *
 * [387] 字符串中的第一个唯一字符
 */

// @lc code=start
func firstUniqChar(s string) int {
	var counts [26]int

	for _, ch := range s {
		counts[int(ch-'a')] += 1
	}

	for i, ch := range s {
		if counts[int(ch-'a')] == 1 {
			return i
		}
	}
	return -1
}

// @lc code=end

