/*
 * @lc app=leetcode.cn id=101 lang=golang
 *
 * [101] 对称二叉树
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
func isSymmetric(root *TreeNode) bool {
	// 轴对称的条件：结构对称、值对称
	if root == nil {
		return true
	}
	return check(root.Left, root.Right)
}
func check(p, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil || p.Val != q.Val {
		return false
	}
	return check(p.Left, q.Right) && check(p.Right, q.Left)
}

// @lc code=end

