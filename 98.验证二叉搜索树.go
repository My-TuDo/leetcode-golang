/*
 * @lc app=leetcode.cn id=98 lang=golang
 *
 * [98] 验证二叉搜索树
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
func isValidBST(root *TreeNode) bool {
	//  初始调用时，给根节点划定一个全系统的最大范围。
	return helper(root, math.MinInt64, math.MaxInt64)
}

func helper(node *TreeNode, lower, upper int) bool {
	if node == nil {
		return true
	}

	if node.Val <= lower || node.Val >= upper {
		return false
	}

	return helper(node.Left, lower, node.Val) && helper(node.Right, node.Val, upper)
}

// @lc code=end

