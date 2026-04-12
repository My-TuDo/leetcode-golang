/*
 * @lc app=leetcode.cn id=739 lang=golang
 *
 * [739] 每日温度
 */

// @lc code=start
func dailyTemperatures(temperatures []int) []int {
	res := make([]int, len(temperatures))
	stack := []int{} // 存储索引

	for i := 0; i < len(temperatures); i++ {
		// 当栈不为空，且当前温度大于栈顶索引对应的温度
		for len(stack) > 0 && temperatures[i] > temperatures[stack[len(stack)-1]] {
			prevIndex := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			res[prevIndex] = i - prevIndex
		}
		stack = append(stack, i) // 入栈
	}
	return res
}

// @lc code=end

