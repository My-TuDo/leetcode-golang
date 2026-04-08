/*
 * @lc app=leetcode.cn id=637 lang=golang
 *
 * [637] 二叉树的层平均值
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
func averageOfLevels(root *TreeNode) []float64 {
	if root == nil {
		return []float64{}
	}

	average := make([]float64, 0, 16) // 存储平均值的数组
	queue := []*TreeNode{root}        // 存储每层节点的数组
	for len(queue) > 0 {
		size := len(queue)
		var sum float64 // 存储子节点val总和
		for i := 0; i < size; i++ {
			node := queue[0]
			queue[0] = nil // 将节点释放
			queue = queue[1:]
			sum += float64(node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		average = append(average, sum/float64(size))
	}
	return average
}

// @lc code=end

