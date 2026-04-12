/*
 * @lc app=leetcode.cn id=239 lang=golang
 *
 * [239] 滑动窗口最大值
 */

// @lc code=start
func maxSlidingWindow(nums []int, k int) []int {
	if len(nums) == 0 {
		return []int{}
	}
	// deque 存储的是 nums 的索引
	deque := []int{}
	res := make([]int, 0, len(nums)-k+1)

	for i := 0; i < len(nums); i++ {
		// 1. 保持队列单调递减：弹出所有比当前值小的队尾元素
		for len(deque) > 0 && nums[i] >= nums[deque[len(deque)-1]] {
			deque = deque[:len(deque)-1]
		}

		// 2. 入队
		deque = append(deque, i)

		// 3. 检查队首是否已经超出窗口范围
		if deque[0] <= i-k {
			deque = deque[1:]
		}
		// 4.当窗口形成后， 记录结果
		if i >= k-1 {
			res = append(res, nums[deque[0]])
		}
	}
	return res
}

// @lc code=end

