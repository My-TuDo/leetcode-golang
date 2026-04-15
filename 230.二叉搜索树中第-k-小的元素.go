/*
 * @lc app=leetcode.cn id=230 lang=golang
 *
 * [230] 二叉搜索树中第 K 小的元素
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
func kthSmallest(root *TreeNode, k int) int {
	// 先深度优先，递归到左子树的最后一个节点，然后在一个一个排序
	var res int
	count := 0

	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		// 终止条件
		if node == nil {
			return
		}

		dfs(node.Left)

		count++
		if count == k {
			res = node.Val
			return
		}

		// 如果还没找到，才往右走
		if count < k {
			dfs(node.Right)
		}
	}

	dfs(root)
	return res
}

// @lc code=end

