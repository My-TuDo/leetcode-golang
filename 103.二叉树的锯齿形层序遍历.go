/*
 * @lc app=leetcode.cn id=103 lang=golang
 *
 * [103] 二叉树的锯齿形层序遍历
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
func zigzagLevelOrder(root *TreeNode) [][]int {
	// 先实现一个基本的BFS遍历代码
	if root == nil {
		return [][]int{}
	}
	queue := []*TreeNode{root}
	res := [][]int{}

	// 定义一个奇偶性计算器
	isReveres := false

	for len(queue) > 0 {
		size := len(queue)
		localNums := make([]int, size) // 临时切片
		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]
			// isReveres为true时倒叙插入
			if isReveres {
				localNums[size-i-1] = node.Val
			} else {
				localNums[i] = node.Val
			}
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		res = append(res, localNums)
		isReveres = !isReveres // 进行锯齿形层序遍历的关键
	}
	return res
}

// @lc code=end

