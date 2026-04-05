/*
 * @lc app=leetcode.cn id=53 lang=golang
 *
 * [53] 最大子数组和
 */

// @lc code=start
func maxSubArray(nums []int) int {
	// 分治法
	return solve(nums, 0, len(nums)-1)
}

func solve(nums []int, left, right int) int {
	if left == right {
		return nums[left]
	}

	mid := (left + right) / 2

	// 递归处理左右两边
	leftMax := solve(nums, left, mid)
	rightMax := solve(nums, mid+1, right)

	// 计算跨越中间的最大和
	crossMax := findCrossMax(nums, left, mid, right)

	// 返回三者中的最大值
	return max(leftMax, max(rightMax, crossMax))
}

func findCrossMax(nums []int, left, mid, right int) int {
	// 从中间向左扫描
	sum := 0
	leftSum := -1 << 31 // 最小值
	for i := mid; i >= left; i-- {
		sum += nums[i]
		if sum > leftSum {
			leftSum = sum
		}
	}

	// 从中间向右扫描
	sum = 0
	rightSum := -1 << 31
	for i := mid + 1; i <= right; i++ {
		sum += nums[i]
		if sum > rightSum {
			rightSum = sum
		}
	}
	return leftSum + rightSum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

