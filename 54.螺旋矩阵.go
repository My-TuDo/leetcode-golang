/*
 * @lc app=leetcode.cn id=54 lang=golang
 *
 * [54] 螺旋矩阵
 */

// @lc code=start
func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return []int{}
	}

	// 定义边界
	top, bottom := 0, len(matrix)-1
	left, right := 0, len(matrix[0])-1

	res := []int{}

	for left <= right && top <= bottom {
		// 从左到右
		for j := left; j <= right; j++ {
			res = append(res, matrix[top][j])
		}
		top++

		// 从上到下
		for i := top; i <= bottom; i++ {
			res = append(res, matrix[i][right])
		}
		right--

		if top <= bottom {
			// 从右到左
			for j := right; j >= left; j-- {
				res = append(res, matrix[bottom][j])
			}
			bottom--
		}

		if left <= right {
			// 从下到上
			for i := bottom; i >= top; i-- {
				res = append(res, matrix[i][left])
			}
			left++
		}
	}

	return res
}

// @lc code=end

