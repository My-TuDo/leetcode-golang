/*
 * @lc app=leetcode.cn id=124 lang=golang
 *
 * [124] 二叉树中的最大路径和
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxPathSum(root *TreeNode) int {
	maxSum := math.MinInt32

	var maxGain func(*TreeNode) int
	maxGain = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		// 计算左右子树的贡献，若为负数则取0
		left := max(maxGain(node.Left), 0)
		right := max(maxGain(node.Right), 0)

		// 更新全局最大值
		currentPrice := node.Val + left + right
		if currentPrice > maxSum {
			maxSum = currentPrice
		}
		return node.Val + max(left, right)
	}

	maxGain(root)
	return maxSum
}

// @lc code=end

