/*
 * @lc app=leetcode.cn id=104 lang=golang
 *
 * [104] 二叉树的最大深度
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
func maxDepth(root *TreeNode) int {
	// 递归求解二叉树的最大深度
	if root == nil {
		return 0 // 空树的深度为 0
	}
	leftDepth := maxDepth(root.Left)      // 计算左子树的深度
	rightDepth := maxDepth(root.Right)    // 计算右子树的深度
	return max(leftDepth, rightDepth) + 1 // 返回左右子树深度的最大值加 1（当前节点）
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

