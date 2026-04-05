/*
 * @lc app=leetcode.cn id=844 lang=golang
 *
 * [844] 比较含退格的字符串
 */

// @lc code=start
func backspaceCompare(s string, t string) bool {
	// 双指针解法： #只影响它左边的数，因此可以从右往左遍历，遇到#记录一次跳过，遇到非#再执行跳过步骤，然后进行两个元素之间的比较
	i, j := len(s)-1, len(t)-1
	skip_S, skip_T := 0, 0 // 记录跳过次数

	for i >= 0 || j >= 0 {
		// 对s进行退格处理
		for i >= 0 {
			if s[i] == '#' {
				skip_S++
				i--
			} else if skip_S > 0 {
				skip_S--
				i--
			} else {
				break // 遇到了要用来比较的字符
			}
		}

		// 对t进行退格处理
		for j >= 0 {
			if t[j] == '#' {
				skip_T++
				j--
			} else if skip_T > 0 {
				skip_T--
				j--
			} else {
				break
			}
		}

		// 比较两个字符串
		if i >= 0 && j >= 0 {
			if s[i] != t[j] {
				return false
			}
		} else if (i >= 0) != (j >= 0) {
			return false // 一个字符遍历完了另一个还没遍历完，说明长度不等
		}
		i--
		j--
	}
	return true
}

// @lc code=end

