/*
 * @lc app=leetcode.cn id=151 lang=golang
 *
 * [151] 反转字符串中的单词
 */

// @lc code=start
func reverseWords(s string) string {
	b := []byte(s)
	n := len(b)

	//原地移除多余空格
	slow := 0
	for fast := 0; fast < n; fast++ {
		if b[fast] != ' ' {
			if slow != 0 {
				b[slow] = ' '
				slow++
			}
			for fast < n && b[fast] != ' ' {
				b[slow] = b[fast]
				slow++
				fast++
			}
		}
	}
	b = b[:slow] // 切掉多余部分
	reverse(b, 0, len(b)-1)
	i := 0
	for j := 0; j <= len(b); j++ {
		// 遇到空格或者达到末尾，说明找到了一个单词
		if j == len(b) || b[j] == ' ' {
			reverse(b, i, j-1)
			i = j + 1
		}
	}
	return string(b)
}

// 辅助函数，反转指定范围的切片
func reverse(b []byte, start, end int) {
	for start < end {
		b[start], b[end] = b[end], b[start]
		start++
		end--
	}
}

// @lc code=end

