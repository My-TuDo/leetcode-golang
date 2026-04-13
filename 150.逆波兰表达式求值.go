/*
 * @lc app=leetcode.cn id=150 lang=golang
 *
 * [150] 逆波兰表达式求值
 */

// @lc code=start
func evalRPN(tokens []string) int {
	stack := []int{} // 模拟栈
	for _, x := range tokens {
		switch x {
		// 如果是运算符，则将栈中两个数出栈
		case "+", "-", "*", "/":
			n2 := stack[len(stack)-1]
			n1 := stack[len(stack)-2]
			stack = stack[:len(stack)-2]

			// 定义一个变量存储计算值
			var res int
			switch x {
			case "+":
				res = n1 + n2
			case "-":
				res = n1 - n2
			case "*":
				res = n1 * n2
			case "/":
				res = n1 / n2
			}
			// 并将res入栈
			stack = append(stack, res)
		// 若为数，则将其转化为整型后入栈
		default:
			val, _ := strconv.Atoi(x)
			stack = append(stack, val)
		}
	}
	// 栈中只剩一个数，即为最终结果
	return stack[0]
}

// @lc code=end

