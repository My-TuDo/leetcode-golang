/*
 * @lc app=leetcode.cn id=695 lang=golang
 *
 * [695] 岛屿的最大面积
 */

// @lc code=start
func maxAreaOfIsland(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}

	rows := len(grid)
	cols := len(grid[0])
	maxArea := 0

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if grid[r][c] == 1 {
				currentArea := dfs(grid, r, c)
				if currentArea > maxArea {
					maxArea = currentArea
				}
			}

		}
	}
	return maxArea
}

// 淹没法，返回值计数
func dfs(grid [][]int, r, c int) int {
	rows := len(grid)
	cols := len(grid[0])

	if r < 0 || c < 0 || r >= rows || c >= cols || grid[r][c] == 0 {
		return 0
	}

	grid[r][c] = 0

	area := 1
	area += dfs(grid, r-1, c)
	area += dfs(grid, r+1, c)
	area += dfs(grid, r, c-1)
	area += dfs(grid, r, c+1)
	return area
}

// @lc code=end

