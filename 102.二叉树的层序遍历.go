/*
 * @lc app=leetcode.cn id=102 lang=golang
 *
 * [102] 二叉树的层序遍历
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
func levelOrder(root *TreeNode) [][]int {
	// 使用队列进行层序遍历
	if root == nil {
		return [][]int{} // 空树返回空列表
	}
	result := [][]int{}
	queue := []*TreeNode{root} // 初始化队列，包含根节点

	for len(queue) > 0 {
		levelSize := len(queue) // 当前层的节点数
		currentLevel := []int{} // 存储当前层的节点值
		for i := 0; i < levelSize; i++ {
			node := queue[0]                              // 获取队列头部节点
			queue = queue[1:]                             // 弹出队列头部节点
			currentLevel = append(currentLevel, node.Val) // 将当前节点值添加到当前层列表

			// 将左右子节点加入队列
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		result = append(result, currentLevel) // 将当前层列表添加到结果中
	}
	return result
}

// @lc code=end

