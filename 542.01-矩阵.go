/*
 * @lc app=leetcode.cn id=542 lang=golang
 *
 * [542] 01 矩阵
 */

// @lc code=start
func updateMatrix(mat [][]int) [][]int {
	rows, cols := len(mat), len(mat[0])
	queue := [][]int{}

	// 1. 初始化：0进队列，1设为极大值
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if mat[r][c] == 0 {
				queue = append(queue, []int{r, c})
			} else {
				mat[r][c] = 1e9 // 用一个很大的数代表未访问
			}
		}
	}

	// 方向数组：上、下、左、右
	dirs := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	// 2. 开始 BFS 扩散
	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		for _, d := range dirs {
			nr, nc := curr[0]+d[0], curr[1]+d[1]
			// 如果邻居在界内，且通过当前点到达邻居的距离更短
			if nr >= 0 && nr < rows && nc >= 0 && nc < cols {
				if mat[nr][nc] > mat[curr[0]][curr[1]]+1 {
					mat[nr][nc] = mat[curr[0]][curr[1]] + 1
					queue = append(queue, []int{nr, nc})
				}
			}
		}
	}
	return mat
}

// @lc code=end

