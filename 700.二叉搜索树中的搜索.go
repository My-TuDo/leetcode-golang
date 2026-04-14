/*
 * @lc app=leetcode.cn id=700 lang=golang
 *
 * [700] 二叉搜索树中的搜索
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
func searchBST(root *TreeNode, val int) *TreeNode {
	// 迭代版本
	for root != nil {
		if root.Val > val {
			root = root.Left
		} else if root.Val < val {
			root = root.Right
		} else {
			break
		}
	}
	return root
}

// @lc code=end

