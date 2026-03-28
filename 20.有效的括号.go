/*
 * @lc app=leetcode.cn id=20 lang=golang
 *
 * [20] 有效的括号
 */

// @lc code=start
func isValid(s string) bool {
	// 先判断 s 的长度是否为偶数，如果不是偶数，直接返回false
	if len(s)%2 != 0 {
		return false
	}
	// 定义一个栈来存储左括号
	stack := []rune{}
	// 定义一个映射来存储括号的对应关系
	pairs := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}

	for _, c := range s {
		// 如果c是右括号，则进入判断
		if left, isRightBracket := pairs[c]; isRightBracket {
			// 如果栈为空，或者栈顶没有对应的与右括号组合的左括号，则返回false
			if len(stack) == 0 || stack[len(stack)-1] != left {
				return false
			}
			// 有，就出栈
			stack = stack[:len(stack)-1]
		} else {
			// 如果是左括号，那么就入栈
			stack = append(stack, c)
		}
	}
	// 判断栈是否为空
	return len(stack) == 0
}

// @lc code=end

