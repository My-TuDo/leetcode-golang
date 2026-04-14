/*
 * @lc app=leetcode.cn id=994 lang=golang
 *
 * [994] 腐烂的橘子
 */

// @lc code=start
func orangesRotting(grid [][]int) int {
	rows := len(grid)
	cols := len(grid[0])
	queue := [][]int{}
	freshCount := 0

	// 1. 统计新鲜句子并把初始腐烂的橘子入队
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if grid[r][c] == 2 {
				queue = append(queue, []int{r, c})
			} else if grid[r][c] == 1 {
				freshCount++
			}
		}
	}

	// 如果一开始就没有新鲜橘子，直接返回0
	if freshCount == 0 {
		return 0
	}

	minutes := 0
	// 方向数组:上下左右
	dirs := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	// 开始 BFS 扩散
	for len(queue) > 0 && freshCount > 0 {
		minutes++
		size := len(queue)
		// 这一层的循环代表这一分钟内所有橘子的扩散
		for i := 0; i < size; i++ {
			curr := queue[0]
			queue = queue[1:]

			for _, d := range dirs {
				nr, nc := curr[0]+d[0], curr[1]+d[1] // 扩散
				// 如果坐标合法且是新鲜橘子
				if nr >= 0 && nr < rows && nc >= 0 && nc < cols && grid[nr][nc] == 1 {
					grid[nr][nc] = 2 // 变腐烂
					freshCount--
					queue = append(queue, []int{nr, nc})
				}
			}
		}
	}

	// 3. 检查算法还有新鲜橘子
	if freshCount > 0 {
		return -1
	}
	return minutes
}

// @lc code=end

