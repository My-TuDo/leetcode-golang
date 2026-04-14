/*
 * @lc app=leetcode.cn id=46 lang=golang
 *
 * [46] 全排列
 */

// @lc code=start
func permute(nums []int) [][]int {
	var res [][]int // 最终传出切片
	path := []int{} // 存储每一种结果
	// 使用used 数组来记录哪些数字已经被选过了，提高查找效率
	used := make([]bool, len(nums)) // 记录不同组合的存取

	var backtrack func()
	backtrack = func() {
		// 结束条件： 路径填满了
		if len(path) == len(nums) {
			// 拷贝一份 path, 否则 res 里的切片会被后续修改覆盖
			temp := make([]int, len(path))
			copy(temp, path)
			res = append(res, temp)
			return
		}

		for i := 0; i < len(nums); i++ {
			// 如果这个数字已经在路径里了，跳过
			if used[i] {
				continue
			}

			// 做选择
			used[i] = true
			path = append(path, nums[i])

			// 进入下层决策树
			backtrack()

			// 撤销选择（回溯的核心）
			path = path[:len(path)-1]
			used[i] = false
		}
	}

	backtrack()
	return res
}

// @lc code=end

