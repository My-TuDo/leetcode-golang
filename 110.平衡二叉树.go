/*
 * @lc app=leetcode.cn id=110 lang=golang
 *
 * [110] 平衡二叉树
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
func isBalanced(root *TreeNode) bool {
	return getHeight(root) != -1
}

func getHeight(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left := getHeight(root.Left)
	if left == -1 {
		return -1
	}

	right := getHeight(root.Right)
	if right == -1 {
		return -1
	}

	if abs(left-right) > 1 {
		return -1
	}

	return max(left, right) + 1
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// @lc code=end

