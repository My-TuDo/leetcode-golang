/*
 * @lc app=leetcode.cn id=93 lang=golang
 *
 * [93] 复原 IP 地址
 */

// @lc code=start
func restoreIpAddresses(s string) []string {
	var res []string
	var path []string // 存放分段的字符串

	var backtrack func(start int) // start 记录每段起始索引
	backtrack = func(start int) {
		// 终止条件：已经分成4段
		if len(path) == 4 {
			if start == len(s) {
				res = append(res, strings.Join(path, ".")) // 拼接并记录在结果切片中
			}
			return
		}

		// 尝试切割长度为1，2，3的段
		for i := 1; i <= 3; i++ {
			// 越界检查
			if start+i > len(s) {
				break
			}
			sub := s[start : start+i]
			if !isValid(sub) {
				continue
			}
			// 选择
			path = append(path, sub)
			// 递归
			backtrack(start + i)
			// 回溯
			path = path[:len(path)-1] // 将之前的分段弹出
		}
	}
	backtrack(0)
	return res
}

// 辅助函数： 检测分段合法性
func isValid(s string) bool {
	// 不能有前导0
	if len(s) > 1 && s[0] == '0' {
		return false
	}
	// 2. 数值不能超过255
	val, _ := strconv.Atoi(s)
	return val <= 255
}

// @lc code=end

