/*
 * @lc app=leetcode.cn id=1047 lang=golang
 *
 * [1047] 删除字符串中的所有相邻重复项
 */

// @lc code=start
func removeDuplicates(s string) string {
	// 使用一个栈来存储字符
	stack := make([]rune, 0, len(s))

	for _, char := range s {
		// 如果栈不为空且栈顶元素与当前字符相同，则弹出栈顶元素
		if len(stack) > 0 && stack[len(stack)-1] == char {
			stack = stack[:len(stack)-1]
		} else {
			// 否则将当前字符压入栈中
			stack = append(stack, char)
		}
	}

	// 将栈中的字符拼接成结果字符串
	return string(stack)

}

// @lc code=end

