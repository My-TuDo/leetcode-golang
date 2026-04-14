/*
 * @lc app=leetcode.cn id=199 lang=golang
 *
 * [199] 二叉树的右视图
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
func rightSideView(root *TreeNode) []int {
	// 广度优先搜索BFS
	// 每次区队列的最末尾的节点进行返回
	if root == nil {
		return []int{}
	}

	queue := []*TreeNode{root}
	rightview := []int{}
	for len(queue) > 0 {
		size := len(queue)
		rightview = append(rightview, queue[size-1].Val)
		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}
	return rightview
}

// @lc code=end

