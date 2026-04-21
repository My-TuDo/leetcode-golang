/*
 * @lc app=leetcode.cn id=394 lang=golang
 *
 * [394] 字符串解码
 */

// @lc code=start
func decodeString(s string) string {
	// 双栈
	numStack := []int{}    // 存放进入括号前的字符串
	strStack := []string{} // 存放进入括号前的倍数
	res := ""              // 当前层级的字符串内容
	multi := 0             // 当前层级的倍数
	for _, char := range s {
		if char >= '0' && char <= '9' {
			// 处理多位数情况
			multi = multi*10 + int(char-'0')
		} else if char == '[' {
			// 碰到左括号，把当前状态压栈
			numStack = append(numStack, multi)
			strStack = append(strStack, res)
			// 重置，开始处理括号内的新内容
			multi = 0
			res = ""
		} else if char == ']' {
			// 碰到右括号，开始读档结算
			curMulti := numStack[len(numStack)-1]
			numStack = numStack[:len(numStack)-1]

			lastStr := strStack[len(strStack)-1]
			strStack = strStack[:len(strStack)-1]

			// 构造括号内的重复部分,拼接到存档的字符串后面
			temp := ""
			for i := 0; i < curMulti; i++ {
				temp += res
			}
			res = lastStr + temp
		} else {
			// 普通字母，累加到当前字符串
			res += string(char)
		}
	}
	return res
}

// @lc code=end

